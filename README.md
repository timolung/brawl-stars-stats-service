# brawl-stars-stats-service
## Instructions to run
### Docker
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main cmd/server/main.go
sam build -s cmd/server
DOCKER_HOST=unix://$HOME/.docker/run/docker.sock sam local start-api -t template-local.yaml
```

## Zip after build to upload to AWS
```
zip build/main.zip build/main  
```

## Test
### Invoke Individual Events
```
DOCKER_HOST=unix://$HOME/.docker/run/docker.sock sam local invoke -e events/sample.json  
```