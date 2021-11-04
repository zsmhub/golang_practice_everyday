package mq

import (
    "encoding/json"
    "fmt"
    "github.com/nsqio/go-nsq"
    "golang_practice_everyday/nsq_mq/mq/producer"
    "golang_practice_everyday/nsq_mq/mq/transfer"
)

type MessageConsumer struct {
    LocalAddress string // 服务器地址
}

// StartNsqConsume 启动nsq消费者，以后所有的消费者在这里注册
func (c *MessageConsumer) StartNsqConsumer(ListenAddress, localAddress string) {
    c.LocalAddress = localAddress

    nsqConsumer(transfer.LockCustomer{}.Topic(), "1", c.handleLockCustomer, 100, ListenAddress)
}

func (*MessageConsumer) handleLockCustomer(msg *nsq.Message) error {
    defer recover()

    var dto transfer.LockCustomer
    err := json.Unmarshal(msg.Body, &dto)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    msg.Finish()

    if string(msg.Body) == producer.InitTopicData {
        return nil
    }

    // todo 处理业务逻辑
    fmt.Println("消息处理完毕：" + string(msg.Body))

    return nil
}
