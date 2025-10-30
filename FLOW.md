# End-to-End (E2E) Flow Diagram - Team To-Do List Application

> **Note**: The diagrams below are in Mermaid format. To view them as visual diagrams, use GitHub, VS Code with Mermaid extension, or online tools like mermaid.live

## 📋 Quick ASCII Flow Diagram (Immediate View)

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                           🚀 USER JOURNEY FLOW                                 │
└─────────────────────────────────────────────────────────────────────────────────┘

┌─────────────┐
│ 🚀 Start    │
└──────┬──────┘
       │
       ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ 🔐 Login    │───▶│ ✅ Valid?   │───▶│ 🔑 Create   │
│ Form        │    │ Credentials │    │ Session     │
└─────────────┘    └──────┬──────┘    └──────┬──────┘
       ▲                  │ No                │ Yes
       │                  ▼                  ▼
       │            ┌─────────────┐    ┌─────────────┐
       │            │ ❌ Show     │    │ 👤 User     │
       │            │ Error       │    │ Role?       │
       │            └─────────────┘    └──────┬──────┘
       │                                     │
       └─────────────────────────────────────┘
                                              │
                                              ▼
                                    ┌─────────────────┐
                                    │ 👑 Admin        │
                                    │ Dashboard       │
                                    └─────────┬───────┘
                                              │
                                              ▼
                                    ┌─────────────────┐
                                    │ 🔧 Admin        │
                                    │ Features        │
                                    └─────────┬───────┘
                                              │
                    ┌─────────────────────────┼─────────────────────────┐
                    │                         │                         │
                    ▼                         ▼                         ▼
            ┌─────────────┐          ┌─────────────┐          ┌─────────────┐
            │ 👥 User     │          │ 📝 Todo     │          │ 🔍 Filter   │
            │ Management  │          │ Management  │          │ & Search    │
            └─────────────┘          └─────────────┘          └─────────────┘
                    │                         │                         │
                    ▼                         ▼                         ▼
            ┌─────────────┐          ┌─────────────┐          ┌─────────────┐
            │ ➕ Create   │          │ ➕ Add      │          │ 📊 Filter   │
            │ User        │          │ Todo        │          │ by Status   │
            └─────────────┘          └─────────────┘          └─────────────┘
                    │                         │                         │
                    ▼                         ▼                         ▼
            ┌─────────────┐          ┌─────────────┐          ┌─────────────┐
            │ ✏️ Update   │          │ ✅ Complete │          │ 👤 Filter   │
            │ User        │          │ Todo        │          │ by User     │
            └─────────────┘          └─────────────┘          └─────────────┘
                    │                         │                         │
                    ▼                         ▼                         ▼
            ┌─────────────┐          ┌─────────────┐          ┌─────────────┐
            │ 🗑️ Delete   │          │ 🗑️ Delete   │          │ 📄 Pagination│
            │ User        │          │ Todo        │          │             │
            └─────────────┘          └─────────────┘          └─────────────┘

┌─────────────────────────────────────────────────────────────────────────────────┐
│                           👤 GENERAL USER FLOW                                 │
└─────────────────────────────────────────────────────────────────────────────────┘

┌─────────────┐
│ 👤 User     │
│ Login       │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 📊 User     │
│ Dashboard   │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ 📝 Todo     │
│ Management  │
└──────┬──────┘
       │
       ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ ➕ Add      │    │ ✅ Complete │    │ 🔍 Filter   │
│ Todo        │    │ Todo        │    │ & Search    │
└─────────────┘    └─────────────┘    └─────────────┘
       │                  │                  │
       ▼                  ▼                  ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ 👤 Auto-    │    │ ✅ Validate │    │ 📊 Filter   │
│ assign to   │    │ Ownership   │    │ by Status   │
│ Self        │    │             │    │             │
└─────────────┘    └─────────────┘    └─────────────┘
       │                  │                  │
       ▼                  ▼                  ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ 💾 Save     │    │ ✅ Mark     │    │ 👤 Filter   │
│ Todo        │    │ Complete    │    │ by User     │
└─────────────┘    └─────────────┘    └─────────────┘

┌─────────────────────────────────────────────────────────────────────────────────┐
│                           🔄 API INTERACTION FLOW                              │
└─────────────────────────────────────────────────────────────────────────────────┘

┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ 🖥️ Frontend │───▶│ 🌐 HTTP     │───▶│ 🖥️ Go       │───▶│ 💾 Data     │
│ (index.html)│    │ Request     │    │ Server      │    │ Layer       │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
       ▲                  ▲                  ▲                  │
       │                  │                  │                  ▼
       │                  │                  │            ┌─────────────┐
       │                  │                  │            │ 👤 Users    │
       │                  │                  │            │ Slice       │
       │                  │                  │            └─────────────┘
       │                  │                  │                  │
       │                  │                  │                  ▼
       │                  │                  │            ┌─────────────┐
       │                  │                  │            │ 📝 Todos    │
       │                  │                  │            │ Slice       │
       │                  │                  │            └─────────────┘
       │                  │                  │                  │
       │                  │                  │                  ▼
       │                  │                  │            ┌─────────────┐
       │                  │                  │            │ 🔑 Sessions │
       │                  │                  │            │ Map         │
       │                  │                  │            └─────────────┘
       │                  │                  │
       │                  │                  ▼
       │                  │            ┌─────────────┐
       │                  │            │ 📄 JSON     │
       │                  │            │ Response    │
       │                  │            └─────────────┘
       │                  │                  │
       │                  │                  ▼
       │                  │            ┌─────────────┐
       │                  │            │ 📋 HTTP     │
       │                  │            │ Headers     │
       │                  │            └─────────────┘
       │                  │                  │
       │                  │                  ▼
       │                  │            ┌─────────────┐
       │                  │            │ 📊 Status   │
       │                  │            │ Codes       │
       │                  │            └─────────────┘
       │                  │
       │                  ▼
       │            ┌─────────────┐
       │            │ 🌐 HTTP     │
       │            │ Response    │
       │            └─────────────┘
       │
       ▼
┌─────────────┐
│ 🔄 Update   │
│ UI          │
└─────────────┘

┌─────────────────────────────────────────────────────────────────────────────────┐
│                           🎯 KEY FLOW PATTERNS                                 │
└─────────────────────────────────────────────────────────────────────────────────┘

1. 🔐 AUTHENTICATION PATTERN
   ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
   │ User Login  │───▶│ Validate    │───▶│ Create      │───▶│ Role-based  │
   │             │    │ Credentials │    │ Session     │    │ Dashboard   │
   └─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘

2. 📝 CRUD OPERATIONS PATTERN
   ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
   │ Frontend    │───▶│ Backend     │───▶│ Data        │───▶│ UI Update   │
   │ Validation  │    │ Validation  │    │ Operation   │    │ & Feedback  │
   └─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘

3. 🔄 STATE MANAGEMENT PATTERN
   ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
   │ User Action │───▶│ Update      │───▶│ Reactive    │───▶│ UI Refresh  │
   │             │    │ State       │    │ Updates     │    │             │
   └─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘

4. ❌ ERROR HANDLING PATTERN
   ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
   │ Error       │───▶│ Show Error  │───▶│ User        │───▶│ Retry/      │
   │ Occurs      │    │ Message     │    │ Correction  │    │ Continue    │
   └─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘

**Legend:**
- **🟢 Start/End** - Application entry and exit points
- **🔵 Process** - System operations and validations
- **🟠 Decision** - User choices and system decisions
- **🟣 Admin** - Admin-specific functionality
- **🟢 User** - General user functionality
- **🔴 Error** - Error handling and recovery
```

## 🎯 Complete User Journey Flow

```mermaid
flowchart TD
    Start([🚀 User Opens Application]) --> Login{🔐 Authentication Required?}
    
    Login -->|No Session| LoginForm[📝 Login Form]
    Login -->|Valid Session| Dashboard[📊 Dashboard]
    
    LoginForm --> ValidateCredentials{✅ Credentials Valid?}
    ValidateCredentials -->|Invalid| ShowError[❌ Show Error Message]
    ValidateCredentials -->|Valid| CreateSession[🔑 Create User Session]
    
    ShowError --> LoginForm
    CreateSession --> SetUserRole{👤 User Role?}
    
    SetUserRole -->|Admin| AdminDashboard[👑 Admin Dashboard]
    SetUserRole -->|User| UserDashboard[👤 User Dashboard]
    
    AdminDashboard --> AdminFeatures[🔧 Admin Features]
    UserDashboard --> UserFeatures[📝 User Features]
    
    AdminFeatures --> UserManagement[👥 User Management]
    AdminFeatures --> TodoManagement[📋 Todo Management]
    AdminFeatures --> FilterTodos[🔍 Filter & Search]
    
    UserFeatures --> TodoManagement
    UserFeatures --> FilterTodos
    UserFeatures --> PasswordUpdate[🔒 Update Password]
    
    UserManagement --> CreateUser[➕ Create User]
    UserManagement --> UpdateUser[✏️ Update User]
    UserManagement --> DeleteUser[🗑️ Delete User]
    UserManagement --> SearchUsers[🔍 Search Users]
    
    TodoManagement --> AddTodo[➕ Add Todo]
    TodoManagement --> CompleteTodo[✅ Complete Todo]
    TodoManagement --> DeleteTodo[🗑️ Delete Todo]
    TodoManagement --> ViewTodos[👀 View Todos]
    
    FilterTodos --> FilterByStatus[📊 Filter by Status]
    FilterTodos --> FilterByUser[👤 Filter by User]
    FilterTodos --> Pagination[📄 Pagination]
    
    %% Logout Flow
    AdminDashboard --> Logout[🚪 Logout]
    UserDashboard --> Logout
    Logout --> ClearSession[🧹 Clear Session]
    ClearSession --> LoginForm
    
    %% Styling
    classDef startEnd fill:#e8f5e8,stroke:#2e7d32,stroke-width:3px
    classDef process fill:#e3f2fd,stroke:#1565c0,stroke-width:2px
    classDef decision fill:#fff3e0,stroke:#ef6c00,stroke-width:2px
    classDef admin fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px
    classDef user fill:#e8f5e8,stroke:#388e3c,stroke-width:2px
    classDef error fill:#ffebee,stroke:#c62828,stroke-width:2px
    
    class Start,LoginForm,Dashboard startEnd
    class CreateSession,SetUserRole,ClearSession process
    class Login,ValidateCredentials decision
    class AdminDashboard,AdminFeatures,UserManagement admin
    class UserDashboard,UserFeatures,PasswordUpdate user
    class ShowError error
```

## 🔄 Detailed System Flow

### **1. Authentication Flow**

```mermaid
sequenceDiagram
    participant U as User
    participant F as Frontend (index.html)
    participant B as Backend (main.go)
    participant S as Session Store
    
    U->>F: Open application
    F->>B: GET /me (check session)
    B->>S: Validate session
    S-->>B: Session data
    B-->>F: User info or 401
    
    alt No Valid Session
        F->>U: Show login form
        U->>F: Enter credentials
        F->>B: POST /login
        B->>S: Create session
        S-->>B: Session created
        B-->>F: User data + session cookie
        F->>U: Redirect to dashboard
    else Valid Session
        F->>U: Show dashboard
    end
```

### **2. Admin User Flow**

```mermaid
flowchart LR
    AdminLogin[👑 Admin Login] --> AdminDashboard[📊 Admin Dashboard]
    
    AdminDashboard --> TodoSection[📝 Todo Management]
    AdminDashboard --> UserSection[👥 User Management]
    
    TodoSection --> AddTodo[➕ Add Todo]
    TodoSection --> CompleteTodo[✅ Complete Todo]
    TodoSection --> DeleteTodo[🗑️ Delete Todo]
    TodoSection --> FilterTodos[🔍 Filter Todos]
    
    UserSection --> CreateUser[➕ Create User]
    UserSection --> UpdateUser[✏️ Update User]
    UserSection --> DeleteUser[🗑️ Delete User]
    UserSection --> SearchUsers[🔍 Search Users]
    
    AddTodo --> SelectUser[👤 Select User]
    SelectUser --> SaveTodo[💾 Save Todo]
    SaveTodo --> RefreshList[🔄 Refresh Todo List]
    
    CreateUser --> ValidateUser[✅ Validate User Data]
    ValidateUser --> SaveUser[💾 Save User]
    SaveUser --> RefreshUsers[🔄 Refresh User List]
    
    UpdateUser --> OpenModal[📋 Open Update Modal]
    OpenModal --> ValidateUpdate[✅ Validate Update Data]
    ValidateUpdate --> SaveUpdate[💾 Save Update]
    SaveUpdate --> RefreshUsers
    
    DeleteUser --> ConfirmModal[⚠️ Confirmation Modal]
    ConfirmModal --> ValidateDelete[✅ Validate Delete]
    ValidateDelete --> ExecuteDelete[🗑️ Execute Delete]
    ExecuteDelete --> RefreshUsers
```

### **3. General User Flow**

```mermaid
flowchart LR
    UserLogin[👤 User Login] --> UserDashboard[📊 User Dashboard]
    
    UserDashboard --> TodoSection[📝 Todo Management]
    UserDashboard --> PasswordSection[🔒 Password Management]
    
    TodoSection --> AddTodo[➕ Add Todo]
    TodoSection --> CompleteTodo[✅ Complete Todo]
    TodoSection --> FilterTodos[🔍 Filter Todos]
    
    PasswordSection --> UpdatePassword[🔒 Update Password]
    
    AddTodo --> AutoAssign[👤 Auto-assign to Self]
    AutoAssign --> SaveTodo[💾 Save Todo]
    SaveTodo --> RefreshList[🔄 Refresh Todo List]
    
    CompleteTodo --> ValidateOwnership[✅ Validate Ownership]
    ValidateOwnership --> MarkComplete[✅ Mark Complete]
    MarkComplete --> RefreshList
    
    UpdatePassword --> ValidateCurrent[✅ Validate Current Password]
    ValidateCurrent --> ValidateNew[✅ Validate New Password]
    ValidateNew --> SavePassword[💾 Save Password]
    SavePassword --> ShowSuccess[✅ Show Success Message]
```

## 🔄 API Interaction Flow

### **Complete API Flow Diagram**

```mermaid
graph TB
    subgraph "Frontend (index.html)"
        UI[🖥️ User Interface]
        Auth[🔐 Authentication Logic]
        TodoUI[📝 Todo Management UI]
        UserUI[👥 User Management UI]
    end
    
    subgraph "Backend (main.go)"
        Router[🛣️ Gorilla Mux Router]
        AuthHandler[🔐 Auth Handlers]
        TodoHandler[📝 Todo Handlers]
        UserHandler[👥 User Handlers]
        SessionStore[🗄️ Session Store]
    end
    
    subgraph "Data Layer"
        Users[👤 Users Slice]
        Todos[📝 Todos Slice]
        Sessions[🔑 Sessions Map]
    end
    
    %% Authentication Flow
    UI -->|Login Request| Auth
    Auth -->|POST /login| Router
    Router --> AuthHandler
    AuthHandler -->|Validate| Users
    AuthHandler -->|Create Session| SessionStore
    SessionStore -->|Session Cookie| Auth
    Auth -->|User Data| UI
    
    %% Todo Management Flow
    TodoUI -->|CRUD Operations| Router
    Router -->|Route to Handler| TodoHandler
    TodoHandler -->|Validate Session| SessionStore
    TodoHandler -->|Data Operations| Todos
    TodoHandler -->|Response| TodoUI
    
    %% User Management Flow
    UserUI -->|Admin Operations| Router
    Router -->|Route to Handler| UserHandler
    UserHandler -->|Validate Admin| SessionStore
    UserHandler -->|Data Operations| Users
    UserHandler -->|Response| UserUI
    
    %% Styling
    classDef frontend fill:#e3f2fd,stroke:#1565c0,stroke-width:2px
    classDef backend fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px
    classDef data fill:#e8f5e8,stroke:#388e3c,stroke-width:2px
    
    class UI,Auth,TodoUI,UserUI frontend
    class Router,AuthHandler,TodoHandler,UserHandler,SessionStore backend
    class Users,Todos,Sessions data
```

## 📱 User Interface Flow

### **Page Navigation Flow**

```mermaid
stateDiagram-v2
    [*] --> LoginPage : Application Start
    
    LoginPage --> AdminDashboard : Admin Login Success
    LoginPage --> UserDashboard : User Login Success
    LoginPage --> LoginPage : Login Failed
    
    AdminDashboard --> TodoManagement : Click Todo Section
    AdminDashboard --> UserManagement : Click User Section
    AdminDashboard --> LoginPage : Logout
    
    UserDashboard --> TodoManagement : Click Todo Section
    UserDashboard --> PasswordUpdate : Click Password Section
    UserDashboard --> LoginPage : Logout
    
    TodoManagement --> AddTodoModal : Click Add Todo
    TodoManagement --> UpdateTodoModal : Click Update Todo
    TodoManagement --> DeleteConfirmModal : Click Delete Todo
    TodoManagement --> AdminDashboard : Back to Dashboard
    TodoManagement --> UserDashboard : Back to Dashboard
    
    UserManagement --> CreateUserModal : Click Create User
    UserManagement --> UpdateUserModal : Click Update User
    UserManagement --> DeleteUserModal : Click Delete User
    UserManagement --> AdminDashboard : Back to Dashboard
    
    PasswordUpdate --> UserDashboard : Password Updated
    PasswordUpdate --> UserDashboard : Cancel Update
    
    AddTodoModal --> TodoManagement : Todo Created
    AddTodoModal --> TodoManagement : Cancel Create
    
    UpdateTodoModal --> TodoManagement : Todo Updated
    UpdateTodoModal --> TodoManagement : Cancel Update
    
    DeleteConfirmModal --> TodoManagement : Todo Deleted
    DeleteConfirmModal --> TodoManagement : Cancel Delete
    
    CreateUserModal --> UserManagement : User Created
    CreateUserModal --> UserManagement : Cancel Create
    
    UpdateUserModal --> UserManagement : User Updated
    UpdateUserModal --> UserManagement : Cancel Update
    
    DeleteUserModal --> UserManagement : User Deleted
    DeleteUserModal --> UserManagement : Cancel Delete
```

## 🔄 Data Flow Architecture

### **Complete Data Flow Diagram**

```mermaid
graph TB
    subgraph "Client Side"
        Browser[🌐 Browser]
        LocalStorage[💾 Local Storage]
        SessionStorage[🔑 Session Storage]
    end
    
    subgraph "Network Layer"
        HTTP[🌐 HTTP Requests]
        CORS[🔒 CORS Headers]
        Cookies[🍪 Session Cookies]
    end
    
    subgraph "Server Side"
        Server[🖥️ Go Server]
        Router[🛣️ Gorilla Mux]
        Middleware[⚙️ Middleware]
        Handlers[🔧 Request Handlers]
    end
    
    subgraph "Data Layer"
        Memory[💾 In-Memory Storage]
        Users[👤 Users Slice]
        Todos[📝 Todos Slice]
        Sessions[🔑 Sessions Map]
    end
    
    subgraph "Response Layer"
        JSON[📄 JSON Response]
        Headers[📋 HTTP Headers]
        Status[📊 Status Codes]
    end
    
    %% Data Flow
    Browser -->|User Action| HTTP
    HTTP -->|Request| CORS
    CORS -->|Add Headers| Cookies
    Cookies -->|Send| Server
    
    Server -->|Route| Router
    Router -->|Process| Middleware
    Middleware -->|Validate| Handlers
    Handlers -->|Query/Update| Memory
    
    Memory -->|Users| Users
    Memory -->|Todos| Todos
    Memory -->|Sessions| Sessions
    
    Users -->|Data| Handlers
    Todos -->|Data| Handlers
    Sessions -->|Data| Handlers
    
    Handlers -->|Response| JSON
    JSON -->|Format| Headers
    Headers -->|Add| Status
    Status -->|Send| HTTP
    
    HTTP -->|Response| Browser
    Browser -->|Update UI| LocalStorage
    Browser -->|Store Session| SessionStorage
    
    %% Styling
    classDef client fill:#e3f2fd,stroke:#1565c0,stroke-width:2px
    classDef network fill:#fff3e0,stroke:#ef6c00,stroke-width:2px
    classDef server fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px
    classDef data fill:#e8f5e8,stroke:#388e3c,stroke-width:2px
    classDef response fill:#fce4ec,stroke:#c2185b,stroke-width:2px
    
    class Browser,LocalStorage,SessionStorage client
    class HTTP,CORS,Cookies network
    class Server,Router,Middleware,Handlers server
    class Memory,Users,Todos,Sessions data
    class JSON,Headers,Status response
```

## 🎯 Key Flow Patterns

### **1. Authentication Pattern**
- **Session-based authentication** with cookie storage
- **Role-based access control** with UI adaptation
- **Automatic session validation** on each request
- **Graceful logout** with session cleanup

### **2. CRUD Operations Pattern**
- **Frontend validation** before API calls
- **Backend validation** for security
- **Real-time UI updates** after operations
- **Error handling** with user feedback

### **3. State Management Pattern**
- **Centralized state** in JavaScript variables
- **Reactive updates** when data changes
- **Filter and pagination** state management
- **Form state** with validation feedback

### **4. Error Handling Pattern**
- **Frontend validation** for immediate feedback
- **Backend validation** for security
- **User-friendly error messages** with field highlighting
- **Graceful degradation** for network issues

## 🔍 Flow Testing Scenarios

### **Happy Path Scenarios**
1. **Admin Login → Create User → Create Todo → Complete Todo → Logout**
2. **User Login → Create Todo → Complete Todo → Update Password → Logout**
3. **Admin Login → Filter Todos → Update User → Delete Todo → Logout**

### **Error Scenarios**
1. **Invalid Login → Show Error → Retry Login**
2. **Network Error → Show Error → Retry Operation**
3. **Validation Error → Show Field Error → Correct Input**
4. **Session Expired → Redirect to Login → Re-authenticate**

### **Edge Cases**
1. **Last Admin Deletion → Show Error → Prevent Deletion**
2. **Empty Todo List → Show Empty State → Add First Todo**
3. **Large Todo List → Pagination → Navigate Pages**
4. **Concurrent Sessions → Handle Multiple Logins**

This comprehensive flow documentation shows the complete end-to-end journey of users through the application, from initial access to final logout, including all system interactions and data flows.
