
FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

COPY ./backend/ /app

COPY ./backend/symbols.toml ./symbols.toml

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/symbols.toml .


EXPOSE 8080

# Run
CMD ["./main"]