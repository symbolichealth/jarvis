package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/jarvis-ai/jarvis"
)

func main() {
	app := &cli.App{
		Name:  "jarvis",
		Usage: "Interact with the Jarvis assistant",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start a new Jarvis session",
				Action: func(c *cli.Context) error {
					j := jarvis.Start()
					c.App.Metadata = map[string]interface{}{"jarvis": j}
					return nil
				},
			},
			{
				Name:  "chat",
				Usage: "send a chat message",
				Action: func(c *cli.Context) error {
					j, ok := c.App.Metadata["jarvis"].(*jarvis.Jarvis)
					if !ok || j == nil {
						j = jarvis.Start()
						c.App.Metadata["jarvis"] = j
					}
					if c.NArg() == 0 {
						return cli.Exit("message required", 1)
					}
					resp, err := j.Chat(c.Args().First())
					if err != nil {
						return err
					}
					_, err = os.Stdout.WriteString(resp + "\n")
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
