package producer

import (
    "encoding/json"
    "fmt"
    "github.com/nsqio/go-nsq"
    "golang_practice_everyday/nsq_mq/mq/transfer"
    "time"
)

var producer *nsq.Producer

const InitTopicData = "{}"

func StartNsqProducer(addr string) {
    if producer != nil {
        return
    }
    var err error
    cfg := nsq.NewConfig()
    producer, err = nsq.NewProducer(addr, cfg)
    if nil != err {
        panic("nsq new panic")
        return
    }

    err = producer.Ping()
    if nil != err {
        panic("nsq ping panic")
    }
}

// 初始化topic，解决报错：error querying nsqlookupd (http://0.0.0.0:4161/lookup?topic=test) - got response 404 Not Found "{\"message\":\"TOPIC_NOT_FOUND\"}"
func InitTopic(topic string) {
    if err := producer.Publish(topic, []byte(InitTopicData)); err != nil {
        fmt.Println(err)
    }
}

func Publish(trans transfer.MqFeed) {
    body, err := json.Marshal(trans)
    if err != nil {
        fmt.Println(err)
        return
    }
    topic := trans.Topic()
    err = producer.Publish(topic, body)
    if err != nil {
        fmt.Println(err)
    }
}

func DeferredPublish(trans transfer.MqFeed, delay time.Duration) {
    body, err := json.Marshal(trans)
    if err != nil {
        fmt.Println(err)
        return
    }
    err = producer.DeferredPublish(trans.Topic(), delay, body)
    if err != nil {
        fmt.Println(err)
    }
}
func StopProducer() {
    if producer != nil {
        producer.Stop()
    }
    fmt.Println("stop nsq producer")
}
