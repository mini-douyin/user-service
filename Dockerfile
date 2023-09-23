# Start from the latest golang base image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Command to run the Go application directly (instead of the compiled binary)
CMD ["go", "run", "."]
