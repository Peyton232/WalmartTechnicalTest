FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the cmd/api package
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/api

# Final stage
FROM alpine:latest

# Copy the binary from the builder image
COPY --from=builder /docker-gs-ping /docker-gs-ping
COPY ./.env .

# Expose port 8080
EXPOSE 8080

# Define an entry point
ENTRYPOINT ["/docker-gs-ping"]
