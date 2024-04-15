FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY build/golang_user_api /app/golang_user_api
COPY test/jwtRS256.key /app/jwtRS256.key
COPY test/jwtRS256.key.pub.pem /app/jwtRS256.key.pub.pem
ENV SESSION_PUBLIC_KEY=/app/jwtRS256.key.pub.pem
ENV SESSION_PRIVATE_KEY=/app/jwtRS256.key

ENTRYPOINT ["/app/golang_user_api", "api"]
