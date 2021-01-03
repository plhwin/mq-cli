package conf

import (
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2/consumer"

	"github.com/spf13/viper"
)

var (
	RocketMQ rocketMQ
)

type rocketMQ struct {
	Open        bool
	NameServers string
	Credentials mqCredentials
	Consumer    mqConsumer
	Producer    mqProducer
}

type mqCredentials struct {
	AccessKey string
	SecretKey string
}

type mqConsumer struct {
	Topic     string
	Group     string
	Model     consumer.MessageModel
	FromWhere consumer.ConsumeFromWhere
	Order     bool
}

type mqProducer struct {
	Topic      string
	Group      string
	Retries    int
	VipChannel bool
}

func Init() bool {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config read error: %s \n", err))
	}

	initConf()

	log.Printf("conf.RocketMQ: %+v \n", RocketMQ)

	return true
}

func initConf() {
	RocketMQ = rocketMQ{
		Open:        viper.GetBool("rocketmq.open"),
		NameServers: viper.GetString("rocketmq.nameServers"),
		Credentials: mqCredentials{
			AccessKey: viper.GetString("rocketmq.credentials.accessKey"),
			SecretKey: viper.GetString("rocketmq.credentials.secretKey"),
		},
		Consumer: mqConsumer{
			Topic: viper.GetString("rocketmq.consumer.topic"),
			Group: viper.GetString("rocketmq.consumer.group"),
			Order: viper.GetBool("rocketmq.consumer.order"),
		},
		Producer: mqProducer{
			Topic:      viper.GetString("rocketmq.producer.topic"),
			Group:      viper.GetString("rocketmq.producer.group"),
			Retries:    viper.GetInt("rocketmq.producer.retries"),
			VipChannel: viper.GetBool("rocketmq.producer.vipChannel"),
		},
	}
	switch viper.GetString("rocketmq.consumer.model") {
	case "BroadCasting":
		RocketMQ.Consumer.Model = consumer.BroadCasting
	case "Clustering":
		RocketMQ.Consumer.Model = consumer.Clustering
	default:
		RocketMQ.Consumer.Model = consumer.Clustering
	}
	switch viper.GetString("rocketmq.consumer.fromWhere") {
	case "FirstOffset":
		RocketMQ.Consumer.FromWhere = consumer.ConsumeFromFirstOffset
	case "Timestamp":
		RocketMQ.Consumer.FromWhere = consumer.ConsumeFromTimestamp
	default:
		RocketMQ.Consumer.FromWhere = consumer.ConsumeFromLastOffset
	}

	return
}
