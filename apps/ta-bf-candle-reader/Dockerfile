# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /ta-bf-candle-reader

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ta-bf-candle-reader /ta-bf-candle-reader

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ta-bf-candle-reader"]
