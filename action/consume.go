package action

import (
	"log"
	"plhwin/mq-cli/mq"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/urfave/cli/v2"
)

func Consume(c *cli.Context) (err error) {
	nameSrvAddr := c.String("nameSrvAddr")
	consumerGroup := c.String("consumerGroup")
	topic := c.String("topic")
	accessKey := c.String("accessKey")
	secretKey := c.String("secretKey")
	consumeOrder := c.Bool("consumeOrder")

	var consumerModel consumer.MessageModel
	switch c.String("consumerModel") {
	case "BroadCasting":
		consumerModel = consumer.BroadCasting
	default:
		consumerModel = consumer.Clustering
	}

	var consumeFromWhere consumer.ConsumeFromWhere
	switch c.String("consumeFromWhere") {
	case "FirstOffset":
		consumeFromWhere = consumer.ConsumeFromFirstOffset
	case "Timestamp":
		consumeFromWhere = consumer.ConsumeFromTimestamp
	default:
		consumeFromWhere = consumer.ConsumeFromLastOffset
	}

	if err = mq.NewPushConsumer(nameSrvAddr, consumerGroup, accessKey, secretKey, consumerModel, consumeFromWhere, consumeOrder); err != nil {
		log.Println("rocketmq NewPushConsumer error", err)
		return
	}

	log.Println(nameSrvAddr, consumerGroup, topic, accessKey, secretKey, consumerModel, consumeFromWhere, consumeOrder)
	mq.Subscribe(topic)
	return
}
