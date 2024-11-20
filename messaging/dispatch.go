package messaging

// Dispatcher - interface for sending events
type Dispatcher interface {
	OnTick(agent any, src *Ticker)
	OnMessage(agent any, msg *Message, src *Channel)
	OnTrace(agent any, activity any)
}

var (
	MutedDispatcher = new(mutedDispatcher)
)

type mutedDispatcher struct{}

func (n mutedDispatcher) OnTick(agent any, src *Ticker)                   {}
func (n mutedDispatcher) OnMessage(agent any, msg *Message, src *Channel) {}
func (n mutedDispatcher) OnTrace(agent any, activity any)                 {}
