# Team To-Do List Application

A comprehensive collaborative team todo list application built with Go (Gorilla Mux) backend and modern HTML/CSS/JavaScript frontend. Features role-based access control, advanced filtering, pagination, and a polished UI/UX.

## ✨ Features

### 🔐 Authentication & User Management
- **Secure Login System** - Session-based authentication with role-based access
- **User Creation** - Admin can create new users with validation
- **User Updates** - Admin can update user details via modal interface
- **User Deletion** - Admin can delete users with confirmation and safety checks
- **Password Management** - Users can update their own passwords
- **Username Validation** - Alphanumeric only, max 15 characters, no spaces or special characters

### 📝 Todo Management
- **Create Todos** - Add new todos with user assignment
- **Complete Todos** - Mark todos as completed (role-based restrictions)
- **Delete Todos** - Remove todos (admin only)
- **Advanced Filtering** - Filter by status (All/Pending/Completed) and user
- **Smart Sorting** - Newest todos first, completed todos at bottom
- **Pagination** - Handle large todo lists with 10 items per page
- **Auto-scroll** - Navigate to filter section when changing pages

### 🎨 Modern UI/UX
- **Responsive Design** - Works on desktop and mobile devices
- **Smooth Animations** - Hover effects, transitions, and micro-interactions
- **Modal Dialogs** - User update and deletion confirmations
- **Real-time Validation** - Inline error messages with field highlighting
- **Notification System** - Success/error messages for user feedback
- **Consistent Styling** - Rounded corners, gradients, and modern color scheme

### 🔒 Role-Based Access Control
- **Admin Users** - Full access to all features
- **General Users** - Limited to their own todos and password updates
- **Smart UI** - Interface adapts based on user role
- **API Security** - Backend validates permissions for all operations

## 🚀 Quick Start

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Start the server:**
   ```bash
   go run main.go
   ```

3. **Open your browser:**
   Go to `http://localhost:8080`

4. **Login with demo credentials:**
   - **Admin**: username=`admin`, password=`admin`
   - **Users**: username=`alice`/`bob`/`charlie`, password=`password123`

## 📡 API Endpoints

### Authentication
- `POST /login` - User authentication
- `POST /logout` - User logout
- `GET /me` - Get current user info
- `PUT /update-password` - Update user password

### Todo Management
- `GET /todos` - Get all todos (authenticated)
- `POST /todos` - Add new todo (authenticated)
- `PUT /todos/{id}/complete` - Complete a todo (authenticated)
- `DELETE /todos/{id}` - Delete a todo (authenticated)

### User Management (Admin)
- `GET /admin/users` - Get all users (authenticated)
- `POST /admin/users` - Create new user (admin only)
- `PUT /admin/users/{id}` - Update user (admin only)
- `DELETE /admin/users/{id}` - Delete user (admin only)

## 📁 Project Structure

```
to-do-list/
├── main.go          # Go backend server with all API endpoints
├── index.html       # Complete frontend application (single file)
├── go.mod           # Go module dependencies
├── go.sum           # Go module checksums
└── README.md        # This documentation
```

## 🛠 Technology Stack

- **Backend**: Go with Gorilla Mux router
- **Frontend**: HTML5, CSS3, JavaScript (ES6+)
- **Authentication**: Session-based with Gorilla Sessions
- **Data Storage**: In-memory (Go slices) with 15 sample todos
- **CORS**: Cross-origin resource sharing enabled
- **Security**: Input validation, XSS protection, role-based access

## 👥 User Roles & Permissions

### Admin Users
- ✅ Create, read, update, delete todos
- ✅ Complete any todo
- ✅ Create new users
- ✅ Update user details (username, password, role)
- ✅ Delete users (with safety validation)
- ✅ Filter todos by any user
- ✅ Access all API endpoints

### General Users
- ✅ Create, read, update their own todos
- ✅ Complete only their own todos
- ✅ Update their own password
- ✅ Filter todos by user (with "Mine" option)
- ❌ Cannot delete todos
- ❌ Cannot manage other users
- ❌ Cannot access admin-only endpoints

## 🎯 Key Features in Detail

### Advanced Filtering
- **Status Filter**: All, Pending, Completed
- **User Filter**: All users, specific user, or "Mine" for current user
- **Combined Filtering**: Both filters work together
- **Filter Reset**: Automatically resets on logout

### Pagination System
- **10 Items Per Page**: Configurable pagination
- **Smart Navigation**: Previous/Next buttons with state management
- **Page Info**: Shows current page and total pages
- **Auto-scroll**: Scrolls to filter section when changing pages

### Validation & Security
- **Frontend Validation**: Real-time input validation with error messages
- **Backend Validation**: Server-side validation for security
- **Username Rules**: Alphanumeric only, 1-15 characters, no spaces
- **Password Security**: Required for user creation and updates
- **Role Validation**: Prevents unauthorized access to admin functions
- **Last Admin Protection**: Prevents deleting the last admin user

### User Experience
- **Responsive Design**: Works on all screen sizes
- **Loading States**: Visual feedback during API calls
- **Error Handling**: Graceful error display and recovery
- **Form Validation**: Clear error messages under each field
- **Modal Confirmations**: Safe deletion with confirmation dialogs
- **Notification System**: Success/error feedback for all actions

## 🔧 Development

The application is designed as a single-page application (SPA) with a clean separation between frontend and backend. The frontend (`index.html`) handles all UI interactions and communicates with the Go backend via REST API calls using modern JavaScript fetch API.

### Key Design Decisions
- **Single HTML File**: All frontend code in one file for simplicity
- **Session-based Auth**: Secure authentication without JWT complexity
- **In-memory Storage**: Simple data storage for development/demo purposes
- **Role-based UI**: Interface adapts based on user permissions
- **Progressive Enhancement**: Works without JavaScript for basic functionality

## 🚀 Production Considerations

- **Session Security**: Configure proper session store for production
- **Database Integration**: Replace in-memory storage with persistent database
- **HTTPS**: Enable SSL/TLS for secure communication
- **Environment Variables**: Use environment variables for configuration
- **Logging**: Implement proper logging and monitoring
- **Error Handling**: Add comprehensive error tracking