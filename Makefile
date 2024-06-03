# Default variable values
DOCKER_HOST := DOCKER_HOST=unix://$$HOME/.docker/run/docker.sock

.PHONY: build run

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

run:
	go run main.go

run-sam:
	sam local start-api