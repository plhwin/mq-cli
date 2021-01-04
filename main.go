package main

import (
	"log"
	"os"
	"plhwin/mq-cli/action"

	"github.com/urfave/cli/v2"
)

var defaultNameSrvAddr = "127.0.0.1:49876"

func main() {
	app := &cli.App{
		Usage: "A client for rocketmq",
		Commands: []*cli.Command{
			{
				Name:    "consume",
				Aliases: []string{"c"},
				Usage:   "consume message from broker",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "nameSrvAddr", Aliases: []string{"n"}, DefaultText: defaultNameSrvAddr, Value: defaultNameSrvAddr, Usage: "Name server address list, eg: 192.168.0.1:9876;192.168.0.2:9876"},
					&cli.StringFlag{Name: "consumerGroup", Aliases: []string{"g"}, Required: true},
					&cli.StringFlag{Name: "topic", Aliases: []string{"t"}, Required: true},
					&cli.StringFlag{Name: "consumerModel", Aliases: []string{"m"}, DefaultText: "Clustering", Value: "Clustering", Usage: "Clustering or BroadCasting"},
					&cli.StringFlag{Name: "consumeFromWhere", Aliases: []string{"w"}, DefaultText: "LastOffset", Value: "LastOffset", Usage: "LastOffset, FirstOffset or Timestamp"},
					&cli.StringFlag{Name: "accessKey", Aliases: []string{"a"}},
					&cli.StringFlag{Name: "secretKey", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "consumeOrder", Aliases: []string{"o"}, DefaultText: "false", Value: false, Usage: "consume orderly message"},
				},
				Action: action.Consume,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
