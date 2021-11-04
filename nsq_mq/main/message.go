package main

import (
    "fmt"
    "github.com/judwhite/go-svc"
    "golang_practice_everyday/nsq_mq"
    "golang_practice_everyday/nsq_mq/mq"
    "golang_practice_everyday/nsq_mq/mq/producer"
    "net"
    "os"
    "path/filepath"
    "sync"
    "syscall"
)

type messageProgram struct {
    once         sync.Once
    name         string
    localAddress string
}

func main() {
    p := &messageProgram{name: "message", localAddress: GetLocalIP()}
    if err := svc.Run(p, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
        fmt.Println(err)
    }
}

// svc 服务运行框架 程序启动时执行Init+Start, 服务终止时执行Stop
func (p *messageProgram) Init(env svc.Environment) error {
    if env.IsWindowsService() {
        dir := filepath.Dir(os.Args[0])
        return os.Chdir(dir)
    }
    return nil
}

func (p *messageProgram) Start() error {
    // 启动nsq消息队列
    go func() {
        defer recover()
        producer.StartNsqProducer(nsq_mq.NsqPublishAddress)
        new(mq.MessageConsumer).StartNsqConsumer(nsq_mq.NSQConsumers, p.localAddress)
    }()

    fmt.Println("start %s...", p.name)
    return nil
}

func (p *messageProgram) Stop() error {
    p.once.Do(func() {
        defer producer.StopProducer()
    })
    return nil
}

func GetLocalIP() string {
    var ip string
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ip = ipnet.IP.String()
                return ip
            }
        }
    }
    return ""
}