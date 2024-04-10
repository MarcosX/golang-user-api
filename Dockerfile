FROM golang:1.22-alpine as builder

RUN apk add git

ENV GO111MODULE=on
ENV GOPROXY=direct

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/golang_user_api cmd/main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY --from=builder /app/build/golang_user_api /app/golang_user_api

ENTRYPOINT ["/app/golang_user_api", "api"]
