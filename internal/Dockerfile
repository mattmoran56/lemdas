# Use the official Golang image from the Docker Hub
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the entire source code to the container
COPY . .

RUN ls -l

# Change directory to the location of the main.go file
WORKDIR /app/internal/
# Download dependencies
RUN go mod download


RUN ls -l
WORKDIR /app/internal/webApi
RUN ls -l

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.go

RUN ls -l


# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/internal/webApi/webApi .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./webApi"]
