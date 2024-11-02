# Stage 1: Build the binary
FROM golang:1.23.2-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o algotrivia


# Stage 2: Run the binary in Alpine
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/algotrivia .
RUN chmod +x /root/algotrivia

# Install the file utility for debugging
RUN apk add --no-cache file

# Debug: Check if the binary is there and executable
RUN ls -la /root/algotrivia
RUN file /root/algotrivia

EXPOSE 8080

CMD ["./algotrivia"]