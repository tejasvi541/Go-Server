# Stage 1: Build the Go binary
FROM golang:1.23.0 AS builder

WORKDIR /app

# Cache go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM ubuntu:22.04

WORKDIR /app

# Install necessary libraries
RUN apt-get update && \
    apt-get install -y libc6 && \
    apt-get clean

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the app will run
EXPOSE 8080 

# Run the Go binary
CMD ["./main"]
