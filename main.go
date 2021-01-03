package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var defaultNameServers = "127.0.0.1:49876"

func main() {
	app := &cli.App{
		Usage: "A client for rocketmq",
		Commands: []*cli.Command{
			{
				Name:    "consume",
				Aliases: []string{"c"},
				Usage:   "consume message from broker",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "nameServers", Aliases: []string{"n"}, DefaultText: defaultNameServers},
					&cli.StringFlag{Name: "consumerGroup", Aliases: []string{"g"}, Required: true},
					&cli.StringFlag{Name: "topic", Aliases: []string{"t"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("nameServers:", c.String("nameServers"))
					fmt.Println("consumerGroup:", c.String("consumerGroup"))
					fmt.Println("topic:", c.String("topic"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
