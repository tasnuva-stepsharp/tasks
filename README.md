
# Tasks API

## Setup Instructions

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Start the PostgreSQL database
4. Update the `dsn` variable in `config/db.go` with your database credentials
5. Run `go run cmd/main/main.go` to start the server

## API Endpoints

* `POST /auth/signup`: Create a new user
* `POST /auth/login`: Login a user
* `POST /tasks`: Create a new task
* `GET /tasks`: Get all tasks
* `GET /tasks/:id`: Get a task by ID
* `PUT /tasks/:id`: Update a task
* `DELETE /tasks/:id`: Delete a task

