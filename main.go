// brawl-stars-stats-service/cmd/server/main.go
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/timolung/brawl-stars-stats-service/internal/config"
	"github.com/timolung/brawl-stars-stats-service/internal/routes"
)

var (
	cfg *config.Config
)

func init() {
	// Initialization code, such as setting up logging, configuration, etc.
	log.Println("Lambda function initializing")
	cfg = &config.Config{
		BrawlStarsAPIKey: os.Getenv("BRAWL_STARS_API_KEY"),
	}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Create a new router
	r := routes.NewRouter()

	// Create a new HTTP request object from the API Gateway proxy request
	req, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(recorder, req)

	// Copy the response from the recorder to the API Gateway proxy response
	resp := recorder.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
