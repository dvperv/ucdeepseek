FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api-gateway .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/api-gateway

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/api-gateway .

EXPOSE 3000

CMD ["./api-gateway"]