# Use official Go image
FROM golang:1.20

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum (if exists)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go application
RUN go build -o book-api

# Expose the port the app runs on
EXPOSE 9090

# Run the application
CMD [ "./book-api" ]
