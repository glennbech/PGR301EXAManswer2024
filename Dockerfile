# Start from a Debian-based image with the Go tools installed
FROM golang:1.20 as builder

# Set the working directory outside the GOPATH to enable Go modules
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a Docker multi-stage build to create a lean production image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /root/

# Copy the binary from the builder stage to the production image
COPY --from=builder /app/main .

# Run the binary program produced by `go install`
CMD ["./main"]