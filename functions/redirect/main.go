package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	URI                   string
	HSTS                  bool
	HSTSMaxAge            time.Duration
	HSTSIncludeSubdomains bool
	HSTSPreload           bool
}

// The Configuration file or section of the HSTS header implementation
var DefaultConfig = &Config{
	URI:                   "http://example.org",
	HSTS:                  true,
	HSTSMaxAge:            300 * 24 * time.Hour,
	HSTSIncludeSubdomains: false,
	HSTSPreload:           false,
}

func loadConfig() *Config {
	c := new(Config)
	*c = *DefaultConfig
	c.loadEnvVars()
	return c
}

func (c *Config) loadEnvVars() {
	var err error
	c.URI = os.Getenv("REDIRECT_TO")
	if c.URI == "" {
		fmt.Printf("Failed to get value from %s environment variable. Using default value.\n", "REDIRECT_TO")
	}

	c.HSTS, err = strconv.ParseBool(os.Getenv("HSTS_ENABLED"))
	if err != nil {
		fmt.Printf("Failed to get value from %s environment variable. Using default value.\n", "HSTS_ENABLED")
	}

	HSTSMaxAge, err := strconv.ParseInt(os.Getenv("HTST_MAX_AGE"), 10, 64)
	if err != nil {
		fmt.Printf("Failed to get value from %s environment variable. Using default value.\n", "HTST_MAX_AGE")
	} else {
		c.HSTSMaxAge = time.Duration(HSTSMaxAge)
	}

	c.HSTSIncludeSubdomains, err = strconv.ParseBool(os.Getenv("HSTS_INCLUDE_SUBDOMAINS"))
	if err != nil {
		fmt.Printf("Failed to get value from %s environment variable. Using default value.\n", "HSTS_INCLUDE_SUBDOMAINS")
	}

	c.HSTSPreload, err = strconv.ParseBool(os.Getenv("HSTS_PRELOAD"))
	if err != nil {
		fmt.Printf("Failed to get value from %s environment variable. Using default value.\n", "HSTS_PRELOAD")
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request from: ", request.Headers["X-Forwarded-For"])

	c := loadConfig()
	fmt.Println("Redirecting to: ", c.URI)

	HSTSHeader := "max-age=" + strconv.FormatInt(int64(c.HSTSMaxAge/time.Second), 10)
	if c.HSTSIncludeSubdomains {
		HSTSHeader += "; includeSubDomains"
	}
	if c.HSTSPreload {
		HSTSHeader += "; preload"
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"Location":                  c.URI,
			"Strict-Transport-Security": HSTSHeader,
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
