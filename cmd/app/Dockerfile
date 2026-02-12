# Stage 1: Build the Go App
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
# Initialize module if not already done
RUN go mod tidy
# Build the binary
RUN go build -o main .

# Stage 2: Run the App (Small Image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./main"]
