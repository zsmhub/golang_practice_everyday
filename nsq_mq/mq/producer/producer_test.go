package producer

import (
    "golang_practice_everyday/nsq_mq"
    "golang_practice_everyday/nsq_mq/mq/transfer"
    "testing"
)

func init() {
    StartNsqProducer(nsq_mq.NsqPublishAddress)
}

func TestPublish(t *testing.T) {
    t.Run("同步数据", func(t *testing.T) {
        Publish(transfer.LockCustomer{
            StaffId: 2939079,
        })
    })
}
