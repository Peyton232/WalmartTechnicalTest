# Use the official Golang image as a base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code and go.mod/go.sum files to the container
COPY ../ ./

# Build the Go application
RUN go build -o myapp

# Expose the port your Go application will listen on (change as needed)
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./myapp"]
