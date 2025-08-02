# Blog API Documentation

## Authentication Endpoints

### 1. User Registration (Sign Up)

**Endpoint:** `POST /sign-up`

**Description:** Register a new user account

**Request Body:**

```json
{
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "password": "password123"
}
```

**Response (Success - 201):**

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "user@example.com",
      "firstName": "John",
      "lastName": "Doe",
      "role": "user",
      "userProfile": {},
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    },
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "message": "User registered successfully"
  }
}
```

**Response (Error - 400):**

```json
{
  "error": "Registration failed",
  "message": "user with this email already exists"
}
```

### 2. User Login (Sign In)

**Endpoint:** `POST /sign-in`

**Description:** Authenticate user and get access tokens

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (Success - 200):**

```json
{
  "success": true,
  "data": {
    "user": {
      "id": "507f1f77bcf86cd799439011",
      "email": "user@example.com",
      "firstName": "John",
      "lastName": "Doe",
      "role": "user",
      "userProfile": {},
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    },
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "message": "Login successful"
  }
}
```

**Response (Error - 401):**

```json
{
  "error": "Login failed",
  "message": "invalid email or password"
}
```

### 3. User Logout (Sign Out)

**Endpoint:** `POST /sign-out`

**Description:** Logout user and invalidate tokens

**Request Body:**

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response (Success - 200):**

```json
{
  "success": true,
  "data": {
    "message": "Logout successful"
  }
}
```

**Response (Error - 401):**

```json
{
  "error": "Logout failed",
  "message": "invalid token"
}
```

## Authentication Flow

1. **Registration:** User provides email, first name, last name, and password

   - System validates input
   - Checks if user already exists
   - Hashes password using bcrypt
   - Creates user in database
   - Generates JWT access and refresh tokens
   - Returns user data and tokens

2. **Login:** User provides email and password

   - System validates input
   - Finds user by email
   - Compares password hash
   - Generates new JWT tokens
   - Updates user tokens in database
   - Returns user data and tokens

3. **Logout:** User provides access token
   - System validates token
   - Finds user by token
   - Deletes all user tokens from database
   - Returns success message

## Token Management

- **Access Token:** Valid for 15 minutes, used for API authentication
- **Refresh Token:** Valid for 7 days, used to get new access tokens
- **Token Storage:** Tokens are stored in the user document in MongoDB
- **Security:** Passwords are hashed using bcrypt with default cost

## Error Handling

All endpoints return consistent error responses with:

- `error`: Error type/category
- `message`: Detailed error message
- Appropriate HTTP status codes (400, 401, 500)

## Security Features

- Password hashing with bcrypt
- JWT token-based authentication
- Token expiration and refresh mechanism
- Input validation
- Secure password comparison (prevents timing attacks)

## Clean Architecture Implementation

This API follows Clean Architecture principles with clear separation of concerns:

### Architecture Layers

1. **Domain Layer** (`Domain/`)

   - Entities (User, Token, Profile)
   - DTOs (SignupRequest, SigninRequest, etc.)
   - Business rules and interfaces

2. **Use Case Layer** (`Usecase/`)

   - Business logic implementation
   - Orchestrates domain entities and infrastructure
   - Implements authentication flows

3. **Repository Layer** (`Repository/`)

   - Data access abstraction
   - MongoDB operations
   - Implements repository interfaces

4. **Infrastructure Layer** (`Infrastructure/`)

   - `password_service.go` - Password hashing and comparison
   - `jwt_service.go` - JWT token management
   - `auth_middleware.go` - HTTP authentication middleware

5. **Delivery Layer** (`Delivery/`)
   - `Controller/` - HTTP request/response handling
   - `Router/` - Route definitions and middleware setup

### Authentication Middleware

The API includes authentication middleware that can be used to protect routes:

#### Required Authentication

```go
// Apply to routes that require authentication
router.POST("/sign-out", authMiddleware.AuthMiddleware(), userController.SignOutHandler)
```

#### Optional Authentication

```go
// Apply to routes that can work with or without authentication
router.GET("/posts", authMiddleware.OptionalAuthMiddleware(), postController.GetPosts)
```

#### Using Protected Routes

To access protected routes, include the JWT token in the Authorization header:

```
Authorization: Bearer <your-access-token>
```

#### Example Protected Route

```go
// Protected route example
router.GET("/profile", authMiddleware.AuthMiddleware(), func(c *gin.Context) {
    userID := c.GetString("userID") // Get user ID from context
    // Handle protected route logic
})
```
