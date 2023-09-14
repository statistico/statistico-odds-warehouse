package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func BuildSuccessResponse(status int, payload interface{}) events.APIGatewayProxyResponse {
	res := response{
		Message: "success",
		Data:    payload,
	}

	body, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: status,
	}
}

func BuildNonSuccessResponse(message string, status int, err error) events.APIGatewayProxyResponse {
	res := response{
		Message: message,
		Data: []errorMessage{
			{
				Message: err.Error(),
				Code:    1,
			},
		},
	}

	body, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: status,
	}
}
