package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	User      string `json:"user"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type SessionData struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

var todos []Todo
var users []User
var nextID = 1
var nextUserID = 1
var store *sessions.CookieStore

func init() {
	// Initialize session store once
	store = sessions.NewCookieStore([]byte("todo-app-secret-key-very-long-and-secure"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: false,     // Set to false for debugging
		Secure:   false,     // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
}

// Recovery middleware to prevent server crashes
func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Handler panic recovered: %v\n", r)
				http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

// GET /todos — Return all todos as JSON (with optional user filter)
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Check for user filter query parameter
	userFilter := r.URL.Query().Get("user")

	if userFilter != "" {
		// Filter todos by user
		var filteredTodos []Todo
		for _, todo := range todos {
			if todo.User == userFilter {
				filteredTodos = append(filteredTodos, todo)
			}
		}
		json.NewEncoder(w).Encode(filteredTodos)
	} else {
		// Return all todos
		json.NewEncoder(w).Encode(todos)
	}
}

// POST /todos — Add a new todo
func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTodo.ID = nextID
	nextID++
	newTodo.Completed = false
	newTodo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	todos = append(todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// DELETE /todos/{id} — Delete a todo by ID
// Complete todo by ID
func completeTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid ID"}`, http.StatusBadRequest)
		return
	}

	// Find and complete todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			json.NewEncoder(w).Encode(map[string]string{"message": "Todo completed successfully"})
			return
		}
	}

	http.Error(w, `{"error": "Todo not found"}`, http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

// Authentication middleware
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "todo-session")
		if err != nil {
			http.Error(w, `{"error": "Session error"}`, http.StatusUnauthorized)
			return
		}

		userID, ok := session.Values["user_id"]
		if !ok {
			http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		r.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
		r.Header.Set("X-User-Role", fmt.Sprintf("%v", session.Values["role"]))
		r.Header.Set("X-User-Name", fmt.Sprintf("%v", session.Values["username"]))

		next.ServeHTTP(w, r)
	}
}

// Admin middleware
func adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-User-Role")
		if role != "admin" {
			http.Error(w, "Forbidden - Admin access required", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// Login endpoint
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Find user
	var user *User
	for _, u := range users {
		if u.Username == loginReq.Username && u.Password == loginReq.Password {
			user = &u
			break
		}
	}

	if user == nil {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	// Create session
	session, err := store.Get(r, "todo-session")
	if err != nil {
		http.Error(w, `{"error": "Session error"}`, http.StatusInternalServerError)
		return
	}

	session.Values["user_id"] = user.ID
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, `{"error": "Session error"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Login successful",
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Logout endpoint
func logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	session, err := store.Get(r, "todo-session")
	if err != nil {
		// Continue with logout even if session is invalid
	} else {
		session.Values["user_id"] = nil
		session.Values["username"] = nil
		session.Values["role"] = nil
		session.Options.MaxAge = -1
		session.Save(r, w)
	}

	response := map[string]string{"message": "Logout successful"}
	json.NewEncoder(w).Encode(response)
}

// Get current user info
func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	session, err := store.Get(r, "todo-session")
	if err != nil {
		// If session is invalid, treat as not logged in
		http.Error(w, `{"error": "Not logged in"}`, http.StatusUnauthorized)
		return
	}

	userID, ok := session.Values["user_id"]
	if !ok {
		http.Error(w, `{"error": "Not logged in"}`, http.StatusUnauthorized)
		return
	}

	user := map[string]interface{}{
		"id":       userID,
		"username": session.Values["username"],
		"role":     session.Values["role"],
	}

	json.NewEncoder(w).Encode(user)
}

// Create user (admin only)
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate username
	if len(newUser.Username) == 0 {
		http.Error(w, `{"error": "Username is required"}`, http.StatusBadRequest)
		return
	}
	if len(newUser.Username) > 15 {
		http.Error(w, `{"error": "Username must be 15 characters or less"}`, http.StatusBadRequest)
		return
	}
	if strings.Contains(newUser.Username, " ") {
		http.Error(w, `{"error": "Username cannot contain spaces"}`, http.StatusBadRequest)
		return
	}
	// Check for special characters - only allow alphanumeric characters
	for _, char := range newUser.Username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			http.Error(w, `{"error": "Username can only contain letters and numbers"}`, http.StatusBadRequest)
			return
		}
	}

	// Check if username already exists
	for _, u := range users {
		if u.Username == newUser.Username {
			http.Error(w, `{"error": "Username already exists"}`, http.StatusConflict)
			return
		}
	}

	newUser.ID = nextUserID
	nextUserID++
	newUser.Role = "user" // Default role

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Get all users (admin only)
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Return users without passwords
	var safeUsers []map[string]interface{}
	for _, user := range users {
		safeUser := map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		}
		safeUsers = append(safeUsers, safeUser)
	}

	json.NewEncoder(w).Encode(safeUsers)
}

// Update password endpoint
func updatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var updateReq UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Get current user from session
	session, _ := store.Get(r, "todo-session")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, `{"error": "Not logged in"}`, http.StatusUnauthorized)
		return
	}

	// Find user
	var user *User
	for i, u := range users {
		if u.ID == userID {
			user = &users[i]
			break
		}
	}

	if user == nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	// Verify current password
	if user.Password != updateReq.CurrentPassword {
		http.Error(w, `{"error": "Current password is incorrect"}`, http.StatusBadRequest)
		return
	}

	// Update password
	user.Password = updateReq.NewPassword

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Get user ID from URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	var updateReq UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	// Validate username
	if len(updateReq.Username) == 0 {
		http.Error(w, `{"error": "Username is required"}`, http.StatusBadRequest)
		return
	}
	if len(updateReq.Username) > 15 {
		http.Error(w, `{"error": "Username must be 15 characters or less"}`, http.StatusBadRequest)
		return
	}
	if strings.Contains(updateReq.Username, " ") {
		http.Error(w, `{"error": "Username cannot contain spaces"}`, http.StatusBadRequest)
		return
	}
	// Check for special characters - only allow alphanumeric characters
	for _, char := range updateReq.Username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			http.Error(w, `{"error": "Username can only contain letters and numbers"}`, http.StatusBadRequest)
			return
		}
	}

	// Validate role
	if updateReq.Role != "admin" && updateReq.Role != "user" {
		http.Error(w, `{"error": "Role must be 'admin' or 'user'"}`, http.StatusBadRequest)
		return
	}

	// Find user
	var user *User
	for i, u := range users {
		if u.ID == userID {
			user = &users[i]
			break
		}
	}

	if user == nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	// Check if username already exists (excluding current user)
	for _, u := range users {
		if u.Username == updateReq.Username && u.ID != userID {
			http.Error(w, `{"error": "Username already exists"}`, http.StatusBadRequest)
			return
		}
	}

	// Update user
	user.Username = updateReq.Username
	user.Role = updateReq.Role

	// Only update password if provided
	if updateReq.Password != "" {
		user.Password = updateReq.Password
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Get user ID from URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusBadRequest)
		return
	}

	// Find user index
	userIndex := -1
	var userToDelete *User
	for i, u := range users {
		if u.ID == userID {
			userIndex = i
			userToDelete = &users[i]
			break
		}
	}

	if userIndex == -1 {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	// Check if user is admin
	if userToDelete.Role == "admin" {
		// Count remaining admin users (excluding the one being deleted)
		adminCount := 0
		for i, u := range users {
			if u.Role == "admin" && i != userIndex {
				adminCount++
			}
		}

		// Prevent deletion if this is the last admin
		if adminCount == 0 {
			http.Error(w, `{"error": "Cannot delete the last admin user. At least one admin must remain in the system."}`, http.StatusBadRequest)
			return
		}
	}

	// Get username for logging
	username := users[userIndex].Username

	// Remove user from users slice
	users = append(users[:userIndex], users[userIndex+1:]...)

	// Remove all todos created by this user
	var remainingTodos []Todo
	for _, todo := range todos {
		if todo.User != username {
			remainingTodos = append(remainingTodos, todo)
		}
	}
	todos = remainingTodos

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// Handle preflight OPTIONS requests for CORS
func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Add panic recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Server panic recovered: %v\n", r)
		}
	}()

	// Initialize with sample data for different users
	now := time.Now()
	todos = []Todo{
		{ID: 1, Text: "Review code changes", Completed: false, User: "alice", CreatedAt: now.Add(-2 * time.Hour).Format("2006-01-02 15:04:05")},
		{ID: 2, Text: "Update documentation", Completed: true, User: "bob", CreatedAt: now.Add(-1 * time.Hour).Format("2006-01-02 15:04:05")},
		{ID: 3, Text: "Fix bug in login", Completed: false, User: "alice", CreatedAt: now.Add(-50 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 4, Text: "Deploy to staging", Completed: false, User: "charlie", CreatedAt: now.Add(-45 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 5, Text: "Write unit tests", Completed: true, User: "alice", CreatedAt: now.Add(-40 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 6, Text: "Design new feature", Completed: false, User: "bob", CreatedAt: now.Add(-35 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 7, Text: "Code review for PR #123", Completed: true, User: "charlie", CreatedAt: now.Add(-30 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 8, Text: "Update API documentation", Completed: false, User: "alice", CreatedAt: now.Add(-25 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 9, Text: "Fix database migration", Completed: true, User: "bob", CreatedAt: now.Add(-20 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 10, Text: "Implement user authentication", Completed: false, User: "charlie", CreatedAt: now.Add(-15 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 11, Text: "Optimize database queries", Completed: true, User: "alice", CreatedAt: now.Add(-10 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 12, Text: "Add error handling", Completed: false, User: "bob", CreatedAt: now.Add(-5 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 13, Text: "Update dependencies", Completed: true, User: "charlie", CreatedAt: now.Add(-3 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 14, Text: "Create user interface mockups", Completed: false, User: "alice", CreatedAt: now.Add(-1 * time.Minute).Format("2006-01-02 15:04:05")},
		{ID: 15, Text: "Set up CI/CD pipeline", Completed: true, User: "bob", CreatedAt: now.Add(-30 * time.Second).Format("2006-01-02 15:04:05")},
	}
	nextID = 16

	// Initialize users
	users = []User{
		{ID: 1, Username: "admin", Password: "admin", Role: "admin"},
		{ID: 2, Username: "alice", Password: "password123", Role: "user"},
		{ID: 3, Username: "bob", Password: "password123", Role: "user"},
		{ID: 4, Username: "charlie", Password: "password123", Role: "user"},
	}
	nextUserID = 5

	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/login", recoveryMiddleware(login)).Methods("POST")
	r.HandleFunc("/logout", recoveryMiddleware(logout)).Methods("POST")
	r.HandleFunc("/me", recoveryMiddleware(getCurrentUser)).Methods("GET")
	r.HandleFunc("/update-password", recoveryMiddleware(authMiddleware(updatePassword))).Methods("POST")

	// User management routes
	r.HandleFunc("/admin/users", authMiddleware(getUsers)).Methods("GET")                            // All authenticated users can see users
	r.HandleFunc("/admin/users", authMiddleware(adminMiddleware(createUser))).Methods("POST")        // Only admin can create
	r.HandleFunc("/admin/users/{id}", authMiddleware(adminMiddleware(updateUser))).Methods("PUT")    // Only admin can update
	r.HandleFunc("/admin/users/{id}", authMiddleware(adminMiddleware(deleteUser))).Methods("DELETE") // Only admin can delete

	// Todo routes (authenticated users)
	r.HandleFunc("/todos", recoveryMiddleware(authMiddleware(getTodos))).Methods("GET")
	r.HandleFunc("/todos", recoveryMiddleware(authMiddleware(addTodo))).Methods("POST")
	r.HandleFunc("/todos/{id}", recoveryMiddleware(authMiddleware(deleteTodo))).Methods("DELETE")
	r.HandleFunc("/todos/{id}/complete", recoveryMiddleware(authMiddleware(completeTodo))).Methods("PUT")

	// Handle CORS preflight requests
	r.HandleFunc("/login", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/logout", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/me", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/update-password", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/admin/users", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/admin/users/{id}", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/todos", handleOptions).Methods("OPTIONS")
	r.HandleFunc("/todos/{id}", handleOptions).Methods("OPTIONS")

	// Serve index.html as root
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("🚀 Team To-Do App Server Starting...")
	fmt.Println("📍 Server: http://localhost:8080")
	fmt.Println("👤 Admin: username=admin, password=admin")
	fmt.Println("👤 Users: username=alice/bob/charlie, password=password123")

	log.Fatal(http.ListenAndServe(":8080", r))
}
