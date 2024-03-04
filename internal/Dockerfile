# Use the official Golang image from the Docker Hub
FROM golang:latest AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the container
COPY upload .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set environment variables


# Copy the compiled Go binary from the builder stage to the new stage
COPY --from=builder /app/main ./webApi

# Expose port 8080 to the outside world
EXPOSE 8080:8080

# Command to run the executable
CMD ["./main"]
