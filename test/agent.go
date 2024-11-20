package test

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"time"
)

type agentT struct {
	agentId string
	ch      *messaging.Channel
}

func NewAgent(uri string) messaging.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = messaging.NewEmissaryChannel(true)
	return a
}

func NewAgentWithChannel(uri string, ch *messaging.Channel) messaging.OpsAgent {
	a := new(agentT)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agentT) Uri() string { return t.agentId }
func (t *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	t.ch.C <- m
	//fmt.Printf("test: opsAgent.Message() -> %v\n", m)
}

func (t *agentT) IsFinalized() bool { return t.ch.IsFinalized() }

// Notify - status notifications
func (t *agentT) Notify(status *core.Status) *core.Status {
	fmt.Printf("test: opsAgent.Handle() -> [status:%v]\n", status)
	status.Handled = true
	return status
}

// Trace - activity tracing
func (t *agentT) Trace(agent messaging.Agent, event, activity string) {
	if agent == nil {
		fmt.Printf("test: opsAgent.Trace() -> %v : %v -> %v %v]\n", core.FmtRFC3339Millis(time.Now().UTC()), agent, event, activity)
	} else {
		fmt.Printf("test: opsAgent.Trace() -> %v : %v -> %v %v]\n", core.FmtRFC3339Millis(time.Now().UTC()), agent, event, activity)
	}
}

func (t *agentT) Run() {
	for {
		select {
		case msg := <-t.ch.C:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				t.finalize()
				return
			default:
			}
		default:
		}
	}
}

func (t *agentT) Shutdown() {
	msg := messaging.NewControlMessage(t.Uri(), t.Uri(), messaging.ShutdownEvent)
	t.ch.Enable()
	t.ch.C <- msg
}

func (t *agentT) finalize() {
	t.ch.Close()
	t.ch = nil
}
