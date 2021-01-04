package action

import (
	"log"
	"plhwin/mq-cli/mq"

	"github.com/urfave/cli/v2"
)

func Produce(c *cli.Context) (err error) {
	nameSrvAddr := c.String("nameSrvAddr")
	producerGroup := c.String("producerGroup")
	topic := c.String("topic")
	accessKey := c.String("accessKey")
	secretKey := c.String("secretKey")
	retries := c.Int("retries")
	vipChannel := c.Bool("vipChannel")
	message := c.String("message")
	shardingKey := c.String("shardingKey")

	if err = mq.NewProducer(nameSrvAddr, producerGroup, accessKey, secretKey, retries, vipChannel); err != nil {
		log.Println("rocketmq NewProducer error", err)
		return
	}

	if err = mq.Producer.Start(); err != nil {
		log.Println("rocketmq Producer start error", err)
		return
	}

	log.Println(nameSrvAddr, producerGroup, topic, accessKey, secretKey, retries, vipChannel, message, shardingKey)

	if err = mq.SendOneWay(topic, message, shardingKey); err != nil {
		log.Println("rocketmq send one way error", err)
		return
	}

	log.Println("message send success!", shardingKey, message)

	return
}
