package mq

import (
	"context"
	"os"
	"plhwin/mq-cli/conf"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var PushConsumer rocketmq.PushConsumer
var Producer rocketmq.Producer

// https://github.com/apache/rocketmq-client-go/blob/master/examples/consumer/orderly/main.go
func NewPushConsumer() (err error) {
	nameServers := []string{conf.RocketMQ.NameServers}
	var opts []consumer.Option
	opts = append(opts, consumer.WithNameServer(nameServers))
	// 分布式程序中每个客户端程序的标识必须不同，
	// 为了避免不必要的问题，ClientIP+instance Name的组合建议唯一，除非有意需要共用连接、资源。
	// https://zhuanlan.zhihu.com/p/27397055
	opts = append(opts, consumer.WithInstance(strconv.Itoa(os.Getpid())))
	opts = append(opts, consumer.WithGroupName(conf.RocketMQ.Consumer.Group))
	opts = append(opts, consumer.WithConsumerModel(conf.RocketMQ.Consumer.Model))
	opts = append(opts, consumer.WithConsumeFromWhere(conf.RocketMQ.Consumer.FromWhere))
	opts = append(opts, consumer.WithConsumerOrder(conf.RocketMQ.Consumer.Order))
	if conf.RocketMQ.Credentials.AccessKey != "" || conf.RocketMQ.Credentials.SecretKey != "" {
		opts = append(opts, consumer.WithCredentials(primitive.Credentials{
			AccessKey: conf.RocketMQ.Credentials.AccessKey,
			SecretKey: conf.RocketMQ.Credentials.SecretKey,
		}))
	}
	PushConsumer, err = rocketmq.NewPushConsumer(opts...)
	return
}

func NewProducer() (err error) {
	nameServers := []string{conf.RocketMQ.NameServers}
	var opts []producer.Option
	opts = append(opts, producer.WithNameServer(nameServers))
	// to use multiple producer instances in a process, should be specify a different name for each instance
	opts = append(opts, producer.WithInstanceName("PD_"+strconv.Itoa(os.Getpid())))
	opts = append(opts, producer.WithGroupName(conf.RocketMQ.Producer.Group))
	opts = append(opts, producer.WithVIPChannel(conf.RocketMQ.Producer.VipChannel))

	if conf.RocketMQ.Credentials.AccessKey != "" || conf.RocketMQ.Credentials.SecretKey != "" {
		opts = append(opts, producer.WithCredentials(primitive.Credentials{
			AccessKey: conf.RocketMQ.Credentials.AccessKey,
			SecretKey: conf.RocketMQ.Credentials.SecretKey,
		}))
	}
	if conf.RocketMQ.Producer.Retries > 0 {
		opts = append(opts, producer.WithRetry(conf.RocketMQ.Producer.Retries))
	}

	Producer, err = rocketmq.NewProducer(opts...)
	return
}

func SendOneWay(topic, message, shardingKey string) (err error) {
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(message),
	}
	if shardingKey != "" {
		msg.WithShardingKey(shardingKey)
	}
	err = Producer.SendOneWay(context.Background(), msg)
	return
}
