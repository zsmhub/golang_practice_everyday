package main

import (
    "github.com/nsqio/go-nsq"
    "log"
    "strconv"
    "time"
)

var (
    NSQPublishAddress = "127.0.0.1:4150" // nsqd
    NSQConsumers      = "127.0.0.1:4150" // nsqlookupd，本地只能直连nsqd，无法连接nsqlookupd(4161)
    // 本地开发环境是用 docker 部署的，nsqd和nsqlookupd是两个容器，consumer.ConnectToNSQLookupd()方法内部也会调用consumer.ConnectToNSQD()方法，
    // docker 容器间是用 container_id 互相通讯，故需要配置 nsqd 的 container_id 到 /etc/hosts，就能解决此问题
    NSQMaxInFlight    = 10               // nsqd 最大连接数
    topic             = "test"
    channel           = "1"
)

func main() {
    go startConsumer()
    startProducer()
}

// 生产者
func startProducer() {
    cfg := nsq.NewConfig()

    producer, err := nsq.NewProducer(NSQPublishAddress, cfg)
    if err != nil {
        log.Fatal(err)
    }

    // 发布消息
    var i uint64 = 1
    for {
        if err := producer.Publish(topic, []byte("test message: "+strconv.FormatUint(i, 10))); err != nil {
            log.Fatal("publish error:" + err.Error())
        }

        time.Sleep(time.Second)
        i++

        if i == 20 {
            break
        }
    }
}

// 消费者
func startConsumer() {
    conf := nsq.NewConfig()
    conf.LookupdPollInterval = 1 * time.Second
    conf.MaxInFlight = NSQMaxInFlight

    consumer, err := nsq.NewConsumer(topic, channel, conf)
    if err != nil {
        log.Fatal(err)
    }

    // 设置消息处理函数
    consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        message.Finish() // 收到消息的响应，否则会一直循环这个方法的逻辑
        log.Println("消费信息：" + string(message.Body))
        return nil
    }))

    // 连接
    if err := consumer.ConnectToNSQD(NSQConsumers); err != nil {
    // if err := consumer.ConnectToNSQLookupd(NSQConsumers); err != nil {
        log.Fatal(err)
    }

    <-consumer.StopChan
}