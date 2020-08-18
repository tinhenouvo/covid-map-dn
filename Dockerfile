FROM golang:1.13-alpine as dev

RUN apk add --no-cache make git curl build-base

COPY cmd/api /app/cmd/api
COPY internal/app /app/internal/app
COPY internal/pkg /app/internal/pkg
COPY schema /app/schema
COPY go.mod go.sum /app/

WORKDIR /app

RUN go build -o ./build/api crawlerdatacovid/cmd/api

EXPOSE 3001

ENTRYPOINT ["/app/build/api"]
