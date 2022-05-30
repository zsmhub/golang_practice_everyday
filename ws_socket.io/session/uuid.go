package session

import (
    "github.com/google/uuid"
    "github.com/googollee/go-socket.io/engineio/session"
)

// 分布式 session id
type UUIDGenerator struct {}

var _ session.IDGenerator = &UUIDGenerator{}

func (g *UUIDGenerator) NewID() string {
    return uuid.New().String()
}

