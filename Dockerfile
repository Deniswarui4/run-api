# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev vips-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates vips

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Create storage directory
RUN mkdir -p storage/events storage/tickets/qrcodes storage/tickets/pdfs

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
