FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o auth-service .

FROM alpine:latest

RUN apk --no-cache add sqlite ca-certificates

WORKDIR /app

COPY --from=builder /app/auth-service .

COPY init-db.sh .
COPY create_db.sql .

RUN chmod +x init-db.sh

RUN ./init-db.sh

EXPOSE 8080

CMD ["./auth-service"]