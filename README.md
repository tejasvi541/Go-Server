# Go-Server Project

![Go Logo](https://blog.golang.org/gopher/header.jpg)

## Overview

The **Go-Server** project is a simple web application built with Go, utilizing the Gin framework for routing, and PostgreSQL as the database. This project demonstrates the core principles of building and deploying a RESTful API with JWT authentication, focusing on managing events and user registration.

### Features

- **User Authentication**: Secure login and signup functionality with JWT tokens.
- **Event Management**: Create, retrieve, update, and delete events.
- **User Event Registration**: Register and unregister for events.
- **Dockerized Deployment**: Easy deployment using Docker and Docker Compose.

### Core Technologies

- **Programming Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker

## Endpoints

### Event Routes

- **GET /events**
  - Fetch all events.
- **GET /events/:id**
  - Fetch a specific event by ID.
- **POST /events**
  - Create a new event (requires authentication).
- **PUT /events/:id**
  - Update an event by ID (requires authentication).
- **DELETE /events/:id**
  - Delete an event by ID (requires authentication).
- **POST /register**
  - Register for an event (requires authentication).
- **POST /unregister**
  - Unregister from an event (requires authentication).

### User Routes

- **POST /signup**
  - Create a new user account.
- **POST /login**
  - Authenticate and log in to receive a JWT token.

## Setup and Installation

### Prerequisites

- Docker installed on your machine
- Docker Compose installed

### Running the Application

1. **Clone the Repository**

   ```bash
   git clone https://github.com/tejasvi541/Go-Server.git
   cd Go-Server
   ```

2. **Build and Run the Docker Containers**

   ```bash
   docker-compose up --build
   ```

3. **Access the Application**

   The server will run on `http://localhost:8080`.

   - Use tools like Postman or curl to interact with the endpoints.
   - Ensure to include the JWT token in the `Authorization` header for routes that require authentication.

### Docker Image

The Docker image for this project can be found on [Docker Hub](https://hub.docker.com/repository/docker/tejasvix/go-server).

### Environment Variables

You can configure the following environment variables in the `.env` file:

- `DB_HOST`: Database host (default is `localhost`)
- `DB_PORT`: Database port (default is `5432`)
- `DB_USER`: Database username (default is `postgres`)
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name (default is `events`)
- `JWT_SECRET`: Secret key for JWT tokens

## Contributing

Contributions are welcome! Please fork the repository and create a pull request to contribute to this project.

## License

This project is licensed under the MIT License.

---

Happy coding! 🎉
