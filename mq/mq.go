package mq

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var PushConsumer rocketmq.PushConsumer
var Producer rocketmq.Producer

// https://github.com/apache/rocketmq-client-go/blob/master/examples/consumer/orderly/main.go
func NewPushConsumer(nameSrvAddr, consumerGroup, accessKey, secretKey string, consumerModel consumer.MessageModel, consumeFromWhere consumer.ConsumeFromWhere, consumeOrder bool) (err error) {
	nameServers := strings.Split(nameSrvAddr, ";")
	var opts []consumer.Option
	opts = append(opts, consumer.WithNameServer(nameServers))
	// 分布式程序中每个客户端程序的标识必须不同，
	// 为了避免不必要的问题，ClientIP+instance Name的组合建议唯一，除非有意需要共用连接、资源。
	// https://zhuanlan.zhihu.com/p/27397055
	opts = append(opts, consumer.WithInstance(strconv.Itoa(os.Getpid())))
	opts = append(opts, consumer.WithGroupName(consumerGroup))
	opts = append(opts, consumer.WithConsumerModel(consumerModel))
	opts = append(opts, consumer.WithConsumeFromWhere(consumeFromWhere))
	opts = append(opts, consumer.WithConsumerOrder(consumeOrder))
	if accessKey != "" || secretKey != "" {
		opts = append(opts, consumer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	}
	PushConsumer, err = rocketmq.NewPushConsumer(opts...)
	return
}

func Subscribe(topic string, consumeLog bool) {
	log.Println("subscribe message from rocketmq start...")
	ch := make(chan struct{})
	err := PushConsumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			// 毫秒
			timeNow := time.Now().UnixNano() / int64(time.Millisecond)
			timeDiff := timeNow - msg.BornTimestamp
			if consumeLog {
				log.Println(msg, timeNow, msg.BornTimestamp, timeDiff)
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Println("Subscribe consumer error", err.Error())
	}
	// Note: start after subscribe
	if err = PushConsumer.Start(); err != nil {
		log.Println("Start consumer error", err.Error())
		os.Exit(0)
	}
	<-ch
	if err = PushConsumer.Shutdown(); err != nil {
		log.Println("Shutdown consumer error", err.Error())
	}
}

func NewProducer(nameSrvAddr, producerGroup, accessKey, secretKey string, retries int, vipChannel bool) (err error) {
	nameServers := strings.Split(nameSrvAddr, ";")
	var opts []producer.Option
	opts = append(opts, producer.WithNameServer(nameServers))
	// to use multiple producer instances in a process, should be specify a different name for each instance
	opts = append(opts, producer.WithInstanceName("PD_"+strconv.Itoa(os.Getpid())))
	opts = append(opts, producer.WithGroupName(producerGroup))
	opts = append(opts, producer.WithVIPChannel(vipChannel))

	if accessKey != "" || secretKey != "" {
		opts = append(opts, producer.WithCredentials(primitive.Credentials{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}))
	}
	if retries > 0 {
		opts = append(opts, producer.WithRetry(retries))
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
