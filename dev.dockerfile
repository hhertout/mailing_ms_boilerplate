FROM golang:1.24.1-alpine3.21

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