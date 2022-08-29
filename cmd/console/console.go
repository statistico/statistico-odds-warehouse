package main

import (
	"fmt"
	"github.com/statistico/statistico-odds-warehouse/internal/app/bootstrap"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	q := app.Queue()
	h := app.MarketHandler()
	l := app.Logger

	console := &cli.App{
		Name: "Statistico Odds Warehouse - Command Line Application",
		Commands: []cli.Command{
			{
				Name:        "market:queue-fetch",
				Usage:       "Fetch and parse markets from queue",
				Description: "Fetch and parse markets from queue",
				Before: func(c *cli.Context) error {
					fmt.Println("Fetching markets from queue...")
					return nil
				},
				After: func(c *cli.Context) error {
					fmt.Println("Complete.")
					return nil
				},
				Action: func(c *cli.Context) error {
					for {
						markets := q.ReceiveMarkets()

						if len(markets) == 0 {
							fmt.Println("Queue is empty. Exiting.")
							break
						}

						for m := range markets {
							err := h.Handle(m)

							if err != nil {
								l.Errorf("Error inserting market %q", err)
							}
						}
					}

					return nil
				},
			},
		},
	}

	err := console.Run(os.Args)

	if err != nil {
		fmt.Printf("Error in executing command: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
