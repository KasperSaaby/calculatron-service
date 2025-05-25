# Use a multi-stage build for a smaller final image
FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build your Go application
# CGO_ENABLED=0 is important for creating static binaries, which are ideal for containers
# GOOS=linux ensures it compiles for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd

# Final stage: a minimal base image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Go service listens on (Cloud Run requires it to listen on $PORT)
# Cloud Run automatically sets the PORT environment variable.
# Your Go application should read this variable, e.g., os.Getenv("PORT").
EXPOSE 8080

# Command to run your service
CMD ["./main"]
