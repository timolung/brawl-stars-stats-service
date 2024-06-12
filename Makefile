.PHONY: build run

build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

zip:
	zip main.zip main

run:
	go run main.go

run-sam:
	sam local start-api