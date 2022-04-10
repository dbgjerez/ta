# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src ./


RUN go build -o /ta-candle-store

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ta-candle-store /ta-candle-store

EXPOSE 8080

USER nonroot:nonroot

ENV GIN_MODE=release

ENTRYPOINT ["/ta-candle-store"]
