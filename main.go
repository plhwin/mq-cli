package main

import (
	"log"
	"os"
	"plhwin/mq-cli/action"

	"github.com/urfave/cli/v2"
)

var defaultNameSrvAddr = "127.0.0.1:9876"

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
					&cli.StringFlag{Name: "consumerGroup", Aliases: []string{"g"}, Required: true, Usage: "consumer group, required"},
					&cli.StringFlag{Name: "topic", Aliases: []string{"t"}, Required: true, Usage: "topic, required"},
					&cli.StringFlag{Name: "consumerModel", Aliases: []string{"m"}, DefaultText: "Clustering", Value: "Clustering", Usage: "Clustering or BroadCasting"},
					&cli.StringFlag{Name: "consumeFromWhere", Aliases: []string{"w"}, DefaultText: "LastOffset", Value: "LastOffset", Usage: "LastOffset, FirstOffset or Timestamp"},
					&cli.StringFlag{Name: "accessKey", Aliases: []string{"a"}, Usage: "credential of accessKey, if needed"},
					&cli.StringFlag{Name: "secretKey", Aliases: []string{"s"}, Usage: "credential of secretKey, if needed"},
					&cli.BoolFlag{Name: "consumeOrder", Aliases: []string{"o"}, DefaultText: "false", Value: false, Usage: "consume orderly message"},
					&cli.BoolFlag{Name: "consumeLog", Aliases: []string{"l"}, DefaultText: "true", Value: true, Usage: "print log when consume message"},
					&cli.Int64Flag{Name: "consumeCountDelay", Aliases: []string{"d"}, DefaultText: "5", Value: 5, Usage: "how many seconds to delay printing consume count"},
				},
				Action: action.Consume,
			},
			{
				Name:    "produce",
				Aliases: []string{"p"},
				Usage:   "produce message to broker",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "nameSrvAddr", Aliases: []string{"n"}, DefaultText: defaultNameSrvAddr, Value: defaultNameSrvAddr, Usage: "Name server address list, eg: 192.168.0.1:9876;192.168.0.2:9876"},
					&cli.StringFlag{Name: "producerGroup", Aliases: []string{"g"}, Required: true, Usage: "producer group, required"},
					&cli.StringFlag{Name: "topic", Aliases: []string{"t"}, Required: true, Usage: "topic, required"},
					&cli.StringFlag{Name: "accessKey", Aliases: []string{"a"}, Usage: "credential of accessKey, if needed"},
					&cli.StringFlag{Name: "secretKey", Aliases: []string{"s"}, Usage: "credential of secretKey, if needed"},
					&cli.BoolFlag{Name: "vipChannel", Aliases: []string{"v"}, DefaultText: "false", Value: false, Usage: "vip channel"},
					&cli.IntFlag{Name: "retries", Aliases: []string{"r"}, DefaultText: "0", Value: 0, Usage: "retry times when send message failed"},
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}, Required: true, Usage: "send message"},
					&cli.StringFlag{Name: "shardingKey", Aliases: []string{"k"}, DefaultText: "", Value: "", Usage: "the sharding key for orderly message"},
				},
				Action: action.Produce,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
