# Forum API

A simple and secure forum API built with Golang, Gin, and GORM, using MariaDB as the database and JWT for user authentication. This API allows users to manage threads and posts with role-based access and provides secure endpoints for basic forum operations.

## Features

- **User Authentication**: Signup/Login with JWT
- **Forum Management**: Create, update, delete threads and posts
- **Role-based Authorization**: Only thread/post creators can modify their content
- **RESTful Design**: Follows standard REST practices for easy integration
- **Error Handling**: Detailed HTTP status codes and error messages

## Requirements

- Go 1.16 or later
- MariaDB
- Git
- Go modules enabled

-> **Clone the repository**:
   ```sh
   git clone https://github.com/your-username/forum-api.git
   cd forum-api
```
**Configure environment variables**: Copy .env.example to .env and set your values:
```
SECRET=your_jwt_secret
DB=your_database
```
**Install dependencies** :
```go
go mod tidy
```
**Run the api**:
```go
go run main.go
```
##Database Schema :
```
-- Table for Users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Table for Threads
CREATE TABLE threads (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Table for Posts
CREATE TABLE posts (
    post_id INT PRIMARY KEY AUTO_INCREMENT,
    content TEXT NOT NULL,
    thread_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```
