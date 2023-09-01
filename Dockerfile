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

# Make the port available to the world outside this container
EXPOSE 8080

# Run the binary when the container starts
CMD ["/app/main"]
