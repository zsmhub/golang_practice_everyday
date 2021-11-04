package nsq_mq

// nsq配置项
var (
    NsqPublishAddress = "127.0.0.1:4150" // nsqd
    NSQConsumers      = "127.0.0.1:4150" // nsqlookupd 127.0.0.1:4161
    NSQMaxInFlight    = 10               // nsqd 最大连接数
)
