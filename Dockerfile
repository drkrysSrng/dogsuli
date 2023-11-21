# Debian for go
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go binary
RUN go build -o app

# Run the application
CMD ["./app"]