package messaging

import "github.com/advanced-go/common/core"

// Dispatcher - interface for sending events
type Dispatcher interface {
	OnTick(agent any, src *Ticker)
	OnMessage(agent any, msg *Message, src *Channel)
	OnError(agent any, status *core.Status) *core.Status
}

var (
	MutedDispatcher = new(mutedDispatcher)
)

type mutedDispatcher struct{}

func (n mutedDispatcher) OnTick(agent any, src *Ticker)                       {}
func (n mutedDispatcher) OnMessage(agent any, msg *Message, src *Channel)     {}
func (n mutedDispatcher) OnError(agent any, status *core.Status) *core.Status { return status }
