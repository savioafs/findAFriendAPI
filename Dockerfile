FROM golang:1.21

# Set working directory
WORKDIR /go/src/app

# Copy only the Go modules and dependencies files
COPY go.mod go.sum ./

# Download Go modules dependencies
RUN go mod download


# Expose the port
EXPOSE 8088

# Build the Go app
RUN go build -o main main.go

# Run the executable
CMD ["./main"]
