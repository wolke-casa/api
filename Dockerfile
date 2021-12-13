FROM docker.io/golang:1.17-alpine

LABEL maintainer="dominic@domm.me" \
      description="Wolke API Docker Image"

WORKDIR /opt/wolke/api

COPY . .

ENV GIN_MODE release

RUN go get -d -v ./...
RUN go build -o api cmd/api/main.go

CMD ["./api"]
