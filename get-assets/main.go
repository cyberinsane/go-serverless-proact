package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(getAssets())
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "assets-handler",
		},
	}

	return resp, nil
}

// TODO - Move this to crud.go
var assets = []asset{
	asset{"someId1", "someName1", "www.google.com", 10},
	asset{"someId2", "someName2", "www.google.com", 10},
}

func getAssets() []asset {
	return assets
}

func getAsset(assetId string) asset {
	return asset{"someId", "someName", "www.google.com", 10}
}

type asset struct {
	ID             string `json:"ref"`
	Name           string `json:"name"`
	CoverImagePath string `json:"coverImagePath"`
	Likes          int    `json:"title"`
}

func main() {
	lambda.Start(Handler)
	body, _ := json.Marshal(getAssets())
	fmt.Println("Assets body: ", body)
}
