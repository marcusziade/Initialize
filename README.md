# Github-api Repository

## Overview

This is an application focused on interacting with the GitHub API. The application features API definitions, multiple endpoints, user-related functionalities, and Docker support for streamlined deployment.

## Directory Structure

```plaintext
Github-api
├── Dockerfile
├── api
│   └── api_definitions
├── cmd
│   └── app
│       └── main.go
├── endpoints
│   ├── User.go
│   ├── endpoints.go
│   ├── starred_repos.go
│   └── starred_repos_test.go
├── entrypoint.sh
├── go.mod
├── go.sum
├── internal
│   └── user
│       ├── handler.go
│       └── service.go
├── main.go
├── pkg
│   └── db
│       └── db.go
└── web
    └── static
```

### Key Components

- `Dockerfile`: Handles containerization of the application. Uses a multi-stage build.
- `cmd/app/main.go`: Main entry point of the application.
- `endpoints`: Houses all the API endpoints including user information and starred repositories.
- `internal`: Contains the business logic for user-related operations.
- `pkg/db`: Database-related logic resides here.
- `web/static`: Static files for any web interface that may be a part of the project.

## Building the Docker Image

```bash
docker build -t github-api .
```

## Running the Docker Container

```bash
docker run -p 8080:8080 -e GITHUB_TOKEN=your_actual_github_token github-api
```

## Dependencies

Managed using Go modules (`go.mod` and `go.sum`).

## Contribution

Feel free to fork this repository, submit pull requests, or report issues.

## License

This project is licensed under the MIT License. See `LICENSE` for more details.
