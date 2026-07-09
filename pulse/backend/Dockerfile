# ===== СТРОИМ =====
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем всё (включая исходники)
COPY . .

# Скачиваем зависимости и собираем
RUN go mod download && \
    go mod verify && \
    go build -o pulse-backend ./cmd/api

# ===== ЗАПУСКАЕМ =====
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/pulse-backend .

EXPOSE 8082

CMD ["./pulse-backend"]