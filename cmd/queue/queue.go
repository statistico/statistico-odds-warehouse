package main

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/bootstrap"
)

func main() {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	q := app.Queue()
	h := app.MarketHandler()
	l := app.Logger

	for {
		fmt.Println("Polling queue for messages...")

		markets := q.ReceiveMarkets()

		for m := range markets {
			err := h.Handle(m)

			if err != nil {
				l.Errorf("Error inserting market %q", err)
			}
		}
	}
}
