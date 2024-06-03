.PHONY: build run

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

run-sam:
	sam local start-api --env-vars env.json