# Use the official Golang image to create a build artifact.
# This is the first stage of a multi-stage build.
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
# COPY go.mod go.sum ./
COPY go.mod  ./
RUN go mod download && go mod verify

COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
