package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi"
	"github.com/timolung/brawl-stars-stats-service/internal/routes"
)

var (
	r      *chi.Mux
	rMutex sync.Mutex
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if the router is initialized
	if r == nil {
		rMutex.Lock()
		defer rMutex.Unlock()

		if r == nil {
			// Initialize the router
			r = routes.SetupRoutes()
		}
	}

	// Convert the APIGatewayProxyRequest to http.Request
	httpRequest, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Forward the request to chi router
	respRecorder := httptest.NewRecorder()
	r.ServeHTTP(respRecorder, httpRequest)

	// Return the response from chi router
	return events.APIGatewayProxyResponse{
		StatusCode: respRecorder.Code,
		Body:       respRecorder.Body.String(),
	}, nil
}

func main() {
	// lambda.Start is how Lambda handles passing various events into your code
	lambda.Start(handler)
}
