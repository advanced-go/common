package messaging

import (
	"github.com/advanced-go/common/core"
)

type Notifier interface {
	OnTick(agent any, ticker *Ticker)
	OnMessage(agent any, msg *Message, ch *Channel)
	OnError(agent any, status *core.Status) *core.Status
}

type MutedNotifier struct{}

func (n MutedNotifier) OnTick(agent any, ticker *Ticker)                    {}
func (n MutedNotifier) OnMessage(agent any, msg *Message, ch *Channel)      {}
func (n MutedNotifier) OnError(agent any, status *core.Status) *core.Status { return status }
