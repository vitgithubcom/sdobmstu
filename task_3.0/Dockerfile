
FROM golang:1.21-alpine AS builder


WORKDIR /app


COPY go.mod ./


RUN go mod download


COPY main.go .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /coinmarketcap-api main.go


FROM alpine:latest


RUN apk --no-cache add ca-certificates


COPY --from=builder /coinmarketcap-api /coinmarketcap-api


RUN chmod +x /coinmarketcap-api


CMD ["/coinmarketcap-api"]