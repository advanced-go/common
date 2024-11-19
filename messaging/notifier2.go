package messaging

import (
	"github.com/advanced-go/common/core"
)

// NotifierOLD - interface for sending notifications that various events have occurred
type NotifierOLD interface {
	OnTick(agent any, src *Ticker)
	OnMessage(agent any, msg *Message, src *Channel)
	OnError(agent any, status *core.Status) *core.Status
}

// MutedNotifierOLD - silent notifications
type MutedNotifierOLD struct{}

func (n MutedNotifierOLD) OnTick(agent any, src *Ticker)                       {}
func (n MutedNotifierOLD) OnMessage(agent any, msg *Message, src *Channel)     {}
func (n MutedNotifierOLD) OnError(agent any, status *core.Status) *core.Status { return status }
