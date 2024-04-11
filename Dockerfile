FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY build/golang_user_api /app/golang_user_api

ENTRYPOINT ["/app/golang_user_api", "api"]
