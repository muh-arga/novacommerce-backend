FROM golang:1.26-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /app/bin/novacommerce \
    ./cmd/api

FROM alpine:3.23 AS runtime

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/bin/novacommerce ./novacommerce

USER appuser

EXPOSE 8080

CMD ["./novacommerce"]