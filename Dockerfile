# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]