// brawl-stars-stats-service/cmd/server/main.go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/timolung/brawl-stars-stats-service/internal/config"
	"github.com/timolung/brawl-stars-stats-service/internal/routes"
)

var (
	gorillaLambda *gorillamux.GorillaMuxAdapter
)

func init() {
	// Initialization code, such as setting up logging, configuration, etc.
	log.Println("Lambda function initializing")
	config.Configure()
	gorillaLambda = gorillamux.New(routes.NewRouter())
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := gorillaLambda.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&req))
	return *r.Version1(), err
}

func main() {
	lambda.Start(Handler)
}
