package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/bootstrap"
	"github.com/statistico/statistico-odds-warehouse/lambda/api/aws"
	"strconv"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())
	reader := app.PostgresMarketReader()
	logger := app.Logger

	id, err := strconv.Atoi(request.PathParameters["id"])

	if err != nil {
		logger.Errorf("error fetching markets from reader: %s", err.Error())

		return aws.BuildNonSuccessResponse(
			"fail",
			422,
			errors.New("id provided is not in the correct format"),
		), nil
	}

	q := warehouse.MarketReaderQuery{}

	if request.QueryStringParameters["name"] != "" {
		q.Market = []string{request.QueryStringParameters["name"]}
	}

	if request.QueryStringParameters["exchange"] != "" {
		q.Exchange = []string{request.QueryStringParameters["exchange"]}
	}

	markets, err := reader.MarketsByEventID(uint64(id), &q)

	if err != nil {
		logger.Errorf("error fetching markets from reader: %s", err.Error())

		return aws.BuildNonSuccessResponse(
			"error",
			500,
			errors.New("internal server error"),
		), nil
	}

	return aws.BuildSuccessResponse(200, markets), nil
}

func main() {
	lambda.Start(handler)
}
