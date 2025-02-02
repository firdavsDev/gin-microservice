# Use the official Go image as a build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o gin-microservice

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gin-microservice .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./gin-microservice"]
