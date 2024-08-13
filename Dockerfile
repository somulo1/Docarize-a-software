# Start off with a specific version as a base image
# Name this stage 'builder'
FROM golang:1.20.14-alpine AS builder

# Add metadata labels for the image
LABEL version="1.0"
LABEL description="Ascii Art Web deployed via containerization"
LABEL maintainer="jesseomolo@gmail.com, nyunja.jp@gmail.com, mcomulosammy37@gmail.com"

# Set working directory inside the container
WORKDIR /app

# Copy in the root directory to working directory
COPY . .

# Build the Go application
RUN go build -o dockerize

# Multi-stage build: Start from a minimal base image
FROM alpine:3.18.3

# Install bash in the final image for runtime access
RUN apk update && apk add --no-cache bash

# Set a non-root user (if possible, for security)
RUN adduser -D appuser
USER appuser

# Set the working directory again
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/. .

# Expose the port on which the app runs
EXPOSE 8000

# Use ENTRYPOINT to define the application startup command
ENTRYPOINT ["./dockerize"]
