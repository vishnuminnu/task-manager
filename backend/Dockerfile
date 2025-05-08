# Build stage: Compile the Go application
FROM golang:1.24 AS builder
WORKDIR /app
# Copy dependency files first to leverage caching
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy the entire source code
COPY . .
# Build the Go binary, disabling CGO for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o task-manager-backend

# Production stage: Create a minimal runtime image
FROM alpine:latest
WORKDIR /app
# Copy the compiled binary from the build stage
COPY --from=builder /app/task-manager-backend .
# Expose port 8080 for the backend API
EXPOSE 8080
# Run the binary
CMD ["./task-manager-backend"]