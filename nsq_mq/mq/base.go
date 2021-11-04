package mq

import (
    "github.com/nsqio/go-nsq"
    "golang_practice_everyday/nsq_mq"
    "golang_practice_everyday/nsq_mq/mq/producer"
    "time"
)

// nsqConsumer 消费消息
func nsqConsumer(topic, channel string, handle func(message *nsq.Message) error, concurrency int, listenAddress string) {
    conf := nsq.NewConfig()
    conf.LookupdPollInterval = 1 * time.Second
    conf.MaxInFlight = nsq_mq.NSQMaxInFlight

    consumer, err := nsq.NewConsumer(topic, channel, conf)
    if err != nil {
        panic(err)
        return
    }
    consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)

    producer.InitTopic(topic)

    if listenAddress == nsq_mq.NsqPublishAddress {
        err = consumer.ConnectToNSQD(listenAddress)
    } else {
        err = consumer.ConnectToNSQLookupd(listenAddress)
    }

    if err != nil {
        panic(err)
    }
}
