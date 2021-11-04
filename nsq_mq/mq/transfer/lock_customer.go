package transfer

// 锁定客户
type LockCustomer struct {
    StaffId        uint64
}

var _ MqFeed = LockCustomer{}

func (msg LockCustomer) Topic() string {
    return "lock_customer"
}
