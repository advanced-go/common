package messaging

import (
	"github.com/advanced-go/common/core"
)

type Notifier interface {
	OnTick(agent any, src *Ticker)
	OnMessage(agent any, msg *Message, src *Channel)
	OnError(agent any, status *core.Status) *core.Status
}

type MutedNotifier struct{}

func (n MutedNotifier) OnTick(agent any, src *Ticker)                       {}
func (n MutedNotifier) OnMessage(agent any, msg *Message, src *Channel)     {}
func (n MutedNotifier) OnError(agent any, status *core.Status) *core.Status { return status }
