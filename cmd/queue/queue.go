package main

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/bootstrap"
)

func main() {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	q := app.Queue()
	r := app.MarketRepository()
	l := app.Logger

	for {
		fmt.Println("Polling queue for messages...")

		markets := q.ReceiveMarkets()

		for m := range markets {
			err := r.Insert(m)

			if err != nil {
				l.Errorf("Error inserting market %q", err)
			}
		}
	}
}
