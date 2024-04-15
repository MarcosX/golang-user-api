# golang-user-api
Golang API with foundational functionality for user signup, update, etc

# Run tests

```
go test ./...
```

# Run container locally

```
docker build . -t local/golang_user_api

docker run -p 8080:8080 local/golang_user_api
```

# Run API locally

```
go run cmd/main.go
```

# Build api

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/golang_user_api cmd/main.go
```

# Login with test credentials

```
curl -X POST http://localhost:8080/login \
  -d 'email=user@email.com&password=pass' \
  -H "Content-Type: application/x-www-form-urlencoded"
```