package router

import (
	"fmt"
	"devport/domain/repository"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

const (
	InstanceGin int = iota
)

func NewWebServerFactory(
	instance int,
	port Port,
	ctxTimeout time.Duration,
	db repository.SQL,
	nosqlDb repository.NoSQL,
) (Server, error) {
	switch instance {
	case InstanceGin:
		return NewGinServer(port, ctxTimeout, db, nosqlDb), nil
	default:
		return nil, fmt.Errorf("instance not exist")
	}
}
