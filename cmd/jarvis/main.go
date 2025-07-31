package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/symbolichealth/jarvis"
)

func main() {
	app := &cli.App{
		Name:  "jarvis",
		Usage: "Interact with the Jarvis assistant",
		Commands: []*cli.Command{
			{
				Name:  "chat",
				Usage: "start an interactive chat session",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						return cli.Exit("chat takes no arguments", 1)
					}

					j := jarvis.Start()
					reader := bufio.NewReader(os.Stdin)

					for {
						fmt.Print("> ")
						input, err := reader.ReadString('\n')
						if err == io.EOF {
							fmt.Println("\nExiting chat.")
							return nil
						}
						if err != nil {
							return err
						}
						input = strings.TrimSpace(input)
						if input == "" {
							continue
						}

						resp, err := j.Chat(input)
						if err != nil {
							return err
						}
						fmt.Printf("\nJarvis: %s\n", resp)
					}
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
