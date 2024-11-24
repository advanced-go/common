package messaging

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"time"
)

type TraceDispatcher interface {
	Tracer
}

type traceDispatch struct {
	all bool
	m   map[string]string
}

func (t *traceDispatch) valid(event string) bool {
	if t.all {
		return true
	}
	if _, ok := t.m[event]; ok {
		return true
	}
	return false
}

func (t *traceDispatch) Trace(agent Agent, channel, event, activity string) {
	if !t.valid(event) {
		return
	}
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	if activity == "" {
		fmt.Printf("trace -> %v [%v] [%v] [%v]\n", core.FmtRFC3339Millis(time.Now().UTC()), channel, event, id)
	} else {
		fmt.Printf("trace -> %v [%v] [%v] [%v] [%v]\n", core.FmtRFC3339Millis(time.Now().UTC()), channel, event, id, activity)
	}
}

func NewTraceDispatcher(events []string) TraceDispatcher {
	t := new(traceDispatch)
	if len(events) == 0 {
		t.all = true
	} else {
		t.m = make(map[string]string)
		for _, event := range events {
			t.m[event] = ""
		}
	}
	return t
}
