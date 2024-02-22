FROM golang:1.19-alpine3.15 AS build_base

WORKDIR /tmp/api

RUN apk add git
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod .
RUN go mod download

COPY . .

RUN swag init

RUN go build -o main .

FROM alpine:3.9 

WORKDIR /app

RUN chgrp -R 0 /app && \
    chmod -R g=u /app

COPY --from=build_base /tmp/api/main api
COPY --from=build_base /tmp/api/config config

CMD ["./api"]
