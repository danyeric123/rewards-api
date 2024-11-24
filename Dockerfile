# Use an official Go image as the base
FROM golang:1.23.3-alpine

# Set the working directory
WORKDIR /app

# Copy and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
