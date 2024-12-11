FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

RUN mkdir -p /app/config
COPY config/serviceAccountKey.json /root/config/serviceAccountKey.json

EXPOSE 8000

CMD ["./main"]