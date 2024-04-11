# Use the official Golang image from the Docker Hub
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the entire source code to the container
COPY .. .

RUN ls -l

# Change directory to the location of the main.go file
WORKDIR /app/internal/
# Download dependencies
RUN go mod download


RUN ls -l
WORKDIR /app/internal/auth
RUN ls -l

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN ls -la

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /app

RUN mkdir .temp

# Copy the built executable from the builder stage
COPY --from=builder /app/internal/auth/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
