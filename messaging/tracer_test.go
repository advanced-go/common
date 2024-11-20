package messaging

import (
	"fmt"
)

type agent struct {
	agentId string
	ch      *Channel
}

func NewAgent(uri string, ch *Channel) Agent {
	a := new(agent)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agent) Uri() string        { return t.agentId }
func (t *agent) Message(m *Message) { fmt.Printf("test: opsAgent.Message() -> %v\n", m) }
func (t *agent) IsFinalized() bool  { return t.ch.IsFinalized() }
func (t *agent) Run()               {}
func (t *agent) Shutdown() {
	if t.ch != nil {
		t.ch.Close()
		t.ch = nil
	}
}

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test", NewEmissaryChannel(true))
	DefaultTracer.Trace(nil, "event:shutdown", "agent shutdown")
	fmt.Printf("\n")

	DefaultTracer.Trace(a, "event:shutdown", "agent shutdown")
	fmt.Printf("\n")

	//Output:
	//<nil> : event:shutdown agent shutdown
	//agent:test : event:shutdown agent shutdown

}
