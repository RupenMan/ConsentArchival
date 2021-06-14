package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"github.com/consent-archival/internal/utils"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	responseValidateHeaders := utils.ValidateHeaders(request.Headers);
	if (responseValidateHeaders != nil) {
		jsonErrorResp,_ := json.Marshal(responseValidateHeaders.Body)
		return events.APIGatewayProxyResponse {
			Body: string(jsonErrorResp[:]),
			StatusCode: responseValidateHeaders.StatusCode,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       string(request.Headers["Accept"]),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
