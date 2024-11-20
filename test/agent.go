package test

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"time"
)

type agent struct {
	agentId string
	ch      *messaging.Channel
}

func NewAgent(uri string, ch *messaging.Channel) messaging.OpsAgent {
	a := new(agent)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agent) Uri() string                  { return t.agentId }
func (t *agent) Message(m *messaging.Message) { fmt.Printf("test: opsAgent.Message() -> %v\n", m) }
func (t *agent) IsFinalized() bool            { return t.ch.IsFinalized() }

// func (t *agent) OnTick(agent any, src *messaging.Ticker)                             {}
// func (t *agent) OnMessage(agent any, msg *messaging.Message, src *messaging.Channel) {}
// func (t *agent) OnTrace(agent any, activity any)                                     {}

// Notify - status notifications
func (t *agent) Notify(status *core.Status) *core.Status {
	fmt.Printf("test: opsAgent.Handle() -> [status:%v]\n", status)
	status.Handled = true
	return status
}

// Trace - activity tracing
func (t *agent) Trace(agent any, activity any) {
	name := "<nil>"
	a := messaging.AgentCast(agent)
	if a != nil {
		name = a.Uri()
	}
	fmt.Printf("test: opsAgent.Trace() -> %v : %v -> %v]\n", core.FmtRFC3339Millis(time.Now().UTC()), name, activity)
}

func (t *agent) Run() {}
func (t *agent) Shutdown() {
	if t.ch != nil {
		t.ch.Close()
		t.ch = nil
	}
}
