FROM golang:1.26.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/lab_08

RUN go build -o contacts-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/lab_08/contacts-app .

EXPOSE 3000

CMD ["./contacts-app"]
