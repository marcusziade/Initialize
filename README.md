# Github-api Repository

## Overview

This is an application focused on interacting with the GitHub API. The application features API definitions, multiple endpoints, user-related functionalities, and Docker support for streamlined deployment.

## Directory Structure

```plaintext
Github-api
├── Dockerfile
├── README.md
├── api
│   └── api_definitions
├── cmd
│   └── app
│       └── main.go
├── endpoints
│   ├── User.go
│   ├── endpoints.go
│   ├── starred_repos.go
│   └── starred_repos_test.go
├── entrypoint.sh
├── go.mod
├── go.sum
└── main.go
```

### Key Components

- `Dockerfile`: Handles containerization of the application.
- `endpoints`: Houses all the API endpoints.

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
