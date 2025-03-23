FROM golang:1.24.1-alpine3.21 AS builder

RUN apk update && apk upgrade
RUN apk add --no-cache  \
    sqlite  \
    sqlite-libs  \
    build-base 

WORKDIR /app

COPY . .

RUN mkdir data

RUN go mod download

RUN go mod tidy

RUN go build -o /tmp/api/main ./cmd/api/main.go

FROM alpine

RUN apk update && apk upgrade
RUN apk add --no-cache  \
    sqlite  \
    sqlite-libs  \
    build-base 

WORKDIR /app

COPY --from=builder /tmp/api/main .

CMD ["./main"]