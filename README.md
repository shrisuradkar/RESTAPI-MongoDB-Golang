# RESTAPI-MongoDB-Golang

# User Management System with Gin (Golang) and MongoDB

This repository contains code for a user management system developed in Golang using the Gin web framework. The application utilizes MongoDB as its database to store user information.

## Features

- User Registration: Create new user profiles with details such as first name, last name, email, and phone number.
- User Retrieval: Retrieve user details by ID.
- User Update: Update existing user information (first name, last name, age, gender, email, phone).
- User Deletion: Delete user profiles or disable them.
- User Listing: Retrieve a list of all users.

## Prerequisites

- Go (Golang) installed on your system.
- MongoDB instance accessible or set up locally.

## Setup

1. Clone the repository:

```git clone https://github.com/your-username/user-management-gin-golang-mongodb.git```

2. Navigate to the project directory:
```cd RESTAPI-MongoDB-Golang```

3. Install dependencies:
```go mod tidy```

4. Set up MongoDB:
Ensure MongoDB is installed and running.
Update the MongoDB connection details in the database configuration file i.e. database/connection.

5. Start the application:
```go run main.go```

Access the API endpoints using a tool like Postman or your preferred HTTP client.

## Endpoints
**POST /user:** Register a new user.
**GET /user/:id:** Retrieve a user by ID.
**PUT /user/:id:** Update user information.
**DELETE /user/:id:** Delete or disable a user.
**GET /user:** List all users with filtering options.

## Postman Collection
Please check https://www.postman.com/solar-crescent-599240/workspace/gin-monogdb-golang/collection/15538441-4e14dd28-182d-433d-9ce7-1cb629e503c2?action=share&creator=15538441

