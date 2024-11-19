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
	a := NewAgent("test: Trace()", NewEmissaryChannel(true))
	DefaultTracer.Trace(a, "agent shutdown")

	fmt.Printf("%v", "\n")

	//Output:
	//test: Trace() : agent shutdown

}
