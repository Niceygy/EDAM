# Use the official Golang image for building
FROM golang:1.25 AS builder
# Set working directory
WORKDIR /app
# Copy Go modules and dependencies
COPY go.mod go.sum ./
# RUN go mod download
# Copy source code
COPY . .
# Build the application
# RUN go build -o edam .
RUN CGO_ENABLED=0 GOOS=linux go build -o edam .
# Use a minimal base image for final deployment
FROM alpine:latest
# Set working directory in the container
WORKDIR /root/
# Copy the built binary from the builder stage
COPY --from=builder /app/edam .
COPY . .
LABEL org.opencontainers.image.description="EDDataCollector"
LABEL org.opencontainers.image.authors="Niceygy (Ava Whale)"

RUN chmod 777 ./edam

# Expose the application port
EXPOSE 3696
# Run the application
CMD ["./edam"]