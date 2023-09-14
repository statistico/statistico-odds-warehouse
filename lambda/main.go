package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/bootstrap"
	"github.com/statistico/statistico-odds-warehouse/internal/warehouse/queue"
)

func parseMessage(ms events.SQSMessage) *queue.EventMarket {
	var mk *queue.EventMarket
	err := json.Unmarshal([]byte(ms.Body), &mk)

	if err != nil {
		//logger.Errorf("Unable to marshal message into market struct, %v.", err)
		return nil
	}

	return mk
}

func handle(event events.SQSEvent) {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig(true))

	handler := app.QueueMarketHandler()
	logger := app.Logger

	for _, message := range event.Records {
		mk := parseMessage(message)

		if err := handler.Handle(mk); err != nil {
			logger.Errorf("Error inserting market %q", err)
		}
	}
}

func main() {
	lambda.Start(handle)
}
