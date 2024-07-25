FROM golang:1.21

# set working directory
WORKDIR /go/src/app

# Create a non-root user and switch to it
RUN useradd -m myuser
USER myuser

# Copy only the Go modules and dependencies files
COPY go.mod go.sum ./

# Download Go modules dependencies
RUN go mod download

# Copy the source code
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./internal ./internal

# EXPOSE the port
EXPOSE 8088

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]
