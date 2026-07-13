# ===== СТРОИМ =====
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копируем всё
COPY . .

# Автоматически создаём go.sum и собираем
RUN go mod tidy && \
    go mod download && \
    go build -o pulse-backend ./cmd/api

# ===== ЗАПУСКАЕМ =====
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/pulse-backend .

EXPOSE 8082

CMD ["./pulse-backend"]