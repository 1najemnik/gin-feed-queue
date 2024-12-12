FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

RUN mkdir -p /app/config
COPY config/serviceAccountKey.json /root/config/serviceAccountKey.json
COPY templates /root/templates

EXPOSE 8080

CMD ["./main"]