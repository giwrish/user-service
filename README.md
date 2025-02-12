# user-service

A simple microserive to manage users. Reason is to learn how to write a mcroservice in go.

Technologies: Go, Postgres, Redis, JWT, Docker, AWS

### 1. Functional Requirements

1. **User Registration**
   - Users can register with a valid email and password.
   - System validates email format and enforces password strength requirements.
   - Duplicate registrations with the same email are prevented.

2. **User Login**
   - Users can log in using their registered email and password.
   - Upon successful login, a JWT token is generated and returned to the user.
   - Invalid login attempts are logged and rate-limited to prevent brute force attacks.

3. **JWT Authentication**
   - JWT tokens are signed and include user information (e.g., user ID, role).
   - Tokens have an expiration time and can be refreshed using a refresh token.
   - Secure storage of secret keys for JWT signing.

4. **Role-Based Authorization (RBAC)**
   - Define user roles (e.g., Admin, User).
   - Access to certain endpoints and actions is restricted based on roles.
   - Admins can manage users (create, update, delete).

5. **Password Security**
   - Passwords are hashed using a secure algorithm (e.g., bcrypt) before storage.
   - Password reset functionality via secure token-based mechanism.

6. **Token Expiration and Refresh**
   - Access tokens expire after a defined period.
   - Refresh tokens allow users to obtain new access tokens without re-authenticating.
   - Mechanism to revoke refresh tokens if compromised.

7. **REST API for User Management**
   - Endpoints for user registration, login, logout, and profile management.
   - Admin endpoints for managing users (CRUD operations).
   - Secure API endpoints with authentication and authorization middleware.

8. **Deployment on AWS**
   - Deploy application using Docker containers on AWS ECS.
   - Use AWS RDS for Postgres database.
   - Deploy Redis for session management and caching.
   - Set up CI/CD pipeline for automated deployments.

### 2. Non-Functional Requirements

1. **Performance**
   - API should handle at least 100 concurrent requests.
   - Token generation and validation should be optimized for speed.

2. **Scalability**
   - System should be designed to scale horizontally by adding more instances.
   - Use Kubernetes for container orchestration to manage scaling.

3. **Security**
   - All data transmission should be encrypted using HTTPS.
   - JWT secret keys and database credentials should be securely managed using AWS Secrets Manager.
   - Implement rate limiting to prevent abuse of authentication endpoints

4. **Logging and Monitoring**
   - Implement logging for key events (e.g., login attempts, token generation).
   - Use AWS CloudWatch or a similar tool for monitoring application performance and errors.
