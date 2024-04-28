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
# Compiles the code into a static binary meaning it includes all necessary dependencies within the binary itself with CGO Disabled.
# -a: tells the Go toolchain to rebuild all packages, even if they are up to date.
# -installsuffix: Used with CGO to distinguish between CGO-enabled and CGO-disabled builds
# -ldflags '-extldflags "-static"': Sets the external linker flags to include -static, which instructs the linker to statically link all libraries, including C libraries, into the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bff ./cmd/server

# Start a new stage from scratch for multi-stage builds.
# This is to reduce the finale docker image size and to isolate build dependencies
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the compiled binary from the previous stage
COPY --from=builder /app/bff .

# Expose the port on which the server will run
EXPOSE 8080

# Command to run the application when starting the container
CMD ["./bff"] 
