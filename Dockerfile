FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth_blacklist ./cmd/

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/auth_blacklist .
COPY .env .env 

RUN chmod +x auth_blacklist

EXPOSE 5555

CMD ["./auth_blacklist"]
