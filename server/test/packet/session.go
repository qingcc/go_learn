package main

//test half

import (
	"fmt"
	"sync"
)

//全局的session管理器
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGc(maxLifeTime int64)
}

var provides = make(map[string]Provider)

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide  %q (forgotten import?)", provideName)
	}
	return &Manager{cookieName: cookieName, provider: provider, maxLifeTime: maxLifeTime}, nil
}

var globalSessions *Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}
