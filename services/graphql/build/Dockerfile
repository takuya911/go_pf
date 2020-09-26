FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine AS prod
WORKDIR /app

COPY --from=builder /src/main .
CMD ["./main"]