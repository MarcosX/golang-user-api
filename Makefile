.PHONY: build

test:
	go test -race -cover -failfast -count=1 ./...

build:
	rm -rf build/golang_user_api
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/golang_user_api cmd/main.go

build-container: build
	docker build . -t local/golang_user_api

run-local:
	go run cmd/main.go

run-container:
	docker run --rm -p 8080:8080 local/golang_user_api

run: build-container run-container
