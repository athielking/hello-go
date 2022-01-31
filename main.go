package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//MyEvent is ... Lambda Function Input Object test
type MyEvent struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, httpEvent events.APIGatewayProxyRequest) (string, error) {
	bytes := []byte(httpEvent.Body)
	var evt MyEvent
	err := json.Unmarshal(bytes, &evt)

	if err != nil {
		return "ERR", err
	}

	return fmt.Sprintf("Hello %s", evt.Name), nil
}

func main() {
	lambda.Start(handleRequest)
}
