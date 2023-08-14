# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR .

# Copy go mod and sum files
COPY go.mod go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main

EXPOSE 3000

# Command to run the executable
ENTRYPOINT ["./main"]