# Task Management
Assignment For Task Management

Project API Documentation
=========================

This document provides detailed usage information for the REST API endpoints. The API, built with the Gin framework, supports operations for user authentication, task management, and user-task assignments.

Server Startup:
---------------
- The server runs on port 8080.

Data Models:
------------
1. Task
   - ID (int)
   - Title (string)
   - Description (string)
   - AssignedUser (int)
   - Status (string)
   - DueDate (string)

2. User
   - ID (int)
   - Name (string)
   - Email (string)

3. UserInput (for registration)
   - Name (string)
   - Email (string)
   - Password (string)

Authentication Endpoints:
-------------------------
1. Login
   - Endpoint: POST /login
   - Payload: email (string), password (string)
   - Response: { "token": "string" }
   - Example: curl -X POST http://localhost:8080/login -d "email=user@example.com&password=12345"

2. Register
   - Endpoint: POST /register
   - Payload: JSON { "name": "string", "email": "string", "password": "string" }
   - Response: { "message": "User created successfully" }
   - Example: curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'

Task Management Endpoints:
--------------------------
1. Create Task
   - Endpoint: POST /tasks
   - Payload: JSON { "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" }
   - Response: { "data": { "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" } }
   - Example: curl -X POST http://localhost:8080/tasks -H "Authorization: Bearer YOUR_TOKEN" -d '{"title": "New Task", "description": "Task description here"}'

2. Get All Tasks
   - Endpoint: GET /tasks
   - Response: [{ "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" }]
   - Example: curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer YOUR_TOKEN"

3. Get Task by ID
   - Endpoint: GET /task_by_id/:task_id
   - Parameters: task_id (int)
   - Response: { "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" }
   - Example: curl -X GET http://localhost:8080/task_by_id/1 -H "Authorization: Bearer YOUR_TOKEN"

4. Get Tasks by User ID
   - Endpoint: GET /task_by_users/:userID
   - Parameters: userID (int)
   - Response: [{ "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" }]
   - Example: curl -X GET http://localhost:8080/task_by_users/1 -H "Authorization: Bearer YOUR_TOKEN"

5. Update Task
   - Endpoint: PUT /tasks/
   - Payload: JSON { "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" }
   - Response: { "data": { "id": "int", "title": "string", "description": "string", "assigned_user": "int", "status": "string", "due_date": "string" } }
   - Example: curl -X PUT http://localhost:8080/tasks/ -H "Authorization: Bearer YOUR_TOKEN" -d '{"title": "Updated Title"}'

6. Delete Task
   - Endpoint: DELETE /tasks/:id
   - Parameters: id (int)
   - Response: { "message": "Task deleted successfully" }
   - Example: curl -X DELETE http://localhost:8080/tasks/1 -H "Authorization: Bearer YOUR_TOKEN"

7. Assign Task
   - Endpoint: POST /assign_task
   - Payload: user_id (int), task_id (int)
   - Response: { "message": "Task assigned successfully" }
   - Example: curl -X POST http://localhost:8080/assign_task -H "Authorization: Bearer YOUR_TOKEN" -d "user_id=1&task_id=2"

Please replace "YOUR_TOKEN" with the actual JWT token obtained after successful authentication using the login endpoint.


## Additional Information

### Testing

Run the following command to execute tests:

```
go test ./tests
```

### Docker

To build and run the application in a Docker container, execute:

```
docker pull ghcr.io/rajendro1/task_management-main:main
docker run -p 8080:8080 ghcr.io/rajendro1/task_management-main:main
```
