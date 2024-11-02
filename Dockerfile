# Use an official Golang image as a base
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Final image
FROM debian:buster
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
