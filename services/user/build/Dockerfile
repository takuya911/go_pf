FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

# 本番用
FROM alpine AS prod
WORKDIR /app
RUN apk add --no-cache curl tzdata

RUN apk update && apk add ca-certificates && update-ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /src/main .

CMD ["./main"]