package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getRedirectURI() string {
	uri := os.Getenv("REDIRECT_TO")

	if uri == "" {
		fmt.Print("Envrionment variable \"REDIRECT_TO\" not set.")
		uri = "http://example.org"
	}

	return uri
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request from: ", request.Headers["X-Forwarded-For"])

	uri := getRedirectURI()
	fmt.Println("Redirecting to: ", uri)

	resp := events.APIGatewayProxyResponse{
		StatusCode: 301,
		Headers: map[string]string{
			"Location": uri,
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
