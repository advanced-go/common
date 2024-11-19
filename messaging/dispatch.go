package messaging

import "github.com/advanced-go/common/core"

// Dispatcher - interface for sending events
type Dispatcher interface {
	OnTick(agent any, src *Ticker)
	OnMessage(agent any, msg *Message, src *Channel)
	OnError(agent any, status *core.Status) *core.Status
}

// MutedDispatcher - silent notifications
type MutedDispatcher struct{}

func (n MutedDispatcher) OnTick(agent any, src *Ticker)                       {}
func (n MutedDispatcher) OnMessage(agent any, msg *Message, src *Channel)     {}
func (n MutedDispatcher) OnError(agent any, status *core.Status) *core.Status { return status }
