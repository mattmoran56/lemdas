# Use the official Golang image from the Docker Hub
FROM golang:1.17-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY . .
# Download dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./webApi/main .

# Start a new stage from scratch
FROM alpine:latest

# Set environment variables


# Expose port 8080 to the outside world
EXPOSE 8080:8080

# Command to run the executable
CMD ["./main"]
