FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY src/ .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o commerce-go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/commerce-go .

EXPOSE 8001

CMD ["./commerce-go"]