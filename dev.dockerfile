FROM golang:1.20-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache  \
    sqlite  \
    sqlite-libs  \
    build-base  \
    make

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . .

RUN mkdir data
RUN go mod download


CMD ["make", "watch"]