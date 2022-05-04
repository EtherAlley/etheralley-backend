# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /main


##
## Deploy
##
FROM gcr.io/distroless/base

WORKDIR /

COPY --from=build /main /main

COPY --from=build /app/assets /assets

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]