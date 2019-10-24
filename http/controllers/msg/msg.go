package ws

import "time"

type Message struct {
	MsgId        string
	MsgType      MessageType
	Content      interface{}
	To           string
	TrySendTimes int
	SendTime     time.Time
}

type MessageType int

const (
	Message_ToOne MessageType = iota
	Message_ToGroup
	Message_ToAll
)
