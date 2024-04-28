# Use the official Golang image as base
FROM golang:1.22.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first to help Docker utilize the Docker layer caching mechanism more efficiently
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application's binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o command_service ./cmd/server

# Start a new stage from scratch for multi-stage builds.
# This is to reduce the finale docker image size and to isolate build dependencies
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the compiled binary from the previous stage
COPY --from=builder /app/command_service .

# Expose the port on which the server will run
EXPOSE 8082

# Command to run the application when starting the container
CMD ["./command_service"] 
