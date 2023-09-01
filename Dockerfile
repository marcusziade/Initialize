# Use an official Go runtime as a builder stage
FROM golang:1.21 AS builder

# Set the working directory in the builder stage
WORKDIR /src

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/main .

# Use a minimal image for running the app
FROM alpine:latest

# Copy the compiled Go binary into this lighter image
COPY --from=builder /out/main /app/main

# Copy the entrypoint script into the image
COPY entrypoint.sh /entrypoint.sh

# Make the script executable
RUN chmod +x /entrypoint.sh

# Make the port available to the world outside this container
EXPOSE 8080

# Set the entry point of the container
ENTRYPOINT ["/entrypoint.sh"]
