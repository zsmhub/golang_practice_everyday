package main

import (
    "github.com/nsqio/go-nsq"
    "log"
    "strconv"
    "time"
)

var (
    listenIP = "127.0.0.1:4150"
    topic = "test"
    channel = "1"
)

func main() {
    go startConsumer()
    startProducer()
}

// 生产者
func startProducer() {
    cfg := nsq.NewConfig()

    producer, err := nsq.NewProducer(listenIP, cfg)
    if err != nil {
        log.Fatal(err)
    }

    // 发布消息
    var i uint64 = 1
    for {
        if err := producer.Publish(topic, []byte("test message: " + strconv.FormatUint(i, 10))); err != nil {
            log.Fatal("publish error:" + err.Error())
        }

        time.Sleep(time.Second)
        i++
    }
}

// 消费者
func startConsumer() {
    cfg := nsq.NewConfig()

    consumer, err := nsq.NewConsumer(topic, channel, cfg)
    if err != nil {
        log.Fatal(err)
    }

    // 设置消息处理函数
    consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        log.Println("消费信息：" + string(message.Body))
        return nil
    }))

    // 连接到单例nsqd
    if err := consumer.ConnectToNSQD(listenIP); err != nil {
        log.Fatal(err)
    }

    <-consumer.StopChan
}