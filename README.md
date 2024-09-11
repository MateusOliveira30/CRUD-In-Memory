# CRUD In-Memory API

This is a simple in-memory API for managing users, implemented in Go using the chi framework. The API supports basic CRUD (Create, Read, Update, Delete) operations for users.

## Project Structure

1. main.go: Entry point of the application.
2. api/: Package containing the API implementation, including handlers and data structures.
   
## Installation

1. Clone the Repository

```
git clone https://github.com/MateusOliveira30/CRUD-In-Memory.git
cd CRUD-In-Memory
```

2. Install Dependencies

Make sure you have Go installed and run:

```
go mod tidy
``` 
3. Running the Server

To start the server, execute:

```
go run main.go
```

The server will be available at http://localhost:8080.

## API Endpoints

1. List All Users
   
Method: GET
Endpoint: /api/users
Description: Returns all users.
Example Response:

```
{
    "c11f0bf5-d1de-4782-b625-3efa402d386f": {
        "first_name": "Maria",
        "last_name": "Mariana",
        "biography": "Second user of the application",
        "id": "c11f0bf5-d1de-4782-b625-3efa402d386f"
    }
}
```

2. Get User by ID

Method: GET
Endpoint: /api/users/{id}
Description: Returns a user based on the provided ID.
URL Parameters:

id: The ID of the user you want to retrieve.
Example Response:

```
{
    "first_name": "Maria",
    "last_name": "Mariana",
    "biography": "Second user of the application",
    "id": "c11f0bf5-d1de-4782-b625-3efa402d386f"
}
```

3. Create New User

Method: POST
Endpoint: /api/users
Description: Creates a new user.
Request Body:

```
{
    "first_name": "Maria",
    "last_name": "Mariana",
    "biography": "New user of the application"
}
```

Example Response:

```
{
    "id": "c11f0bf5-d1de-4782-b625-3efa402d386f"
}
```

4. Update Existing User
   
Method: PUT
Endpoint: /api/users/{id}
Description: Updates an existing user based on the provided ID.
URL Parameters:

id: The ID of the user you want to update.
Request Body:

```
{
    "first_name": "Maria",
    "last_name": "Mariana",
    "biography": "Updated biography of Maria Mariana"
}
```

Response: 204 No Content (no body)

5. Delete User
Method: DELETE
Endpoint: /api/users/{id}
Description: Deletes a user based on the provided ID.
URL Parameters:

id: The ID of the user you want to delete.
Response: 204 No Content (no body)

Notes
User data is stored in memory and will be lost when the server is restarted.
The API does not have authentication or authorization, so anyone with access to the server can manipulate the data.
