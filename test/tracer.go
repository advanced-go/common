package test

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"time"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent messaging.Agent, channel, event, activity string) {
	trace(agent, channel, event, activity)
}

func trace(agent messaging.Agent, channel, event, activity string) {
	id := "<nil>"
	if agent != nil {
		id = agent.Uri()
	}
	if activity == "" {
		fmt.Printf("test: Trace() -> %v [%v] [%v] [%v]\n", core.FmtRFC3339Millis(time.Now().UTC()), channel, event, id)
	} else {
		fmt.Printf("test: Trace() -> %v [%v] [%v] [%v] [%v]\n", core.FmtRFC3339Millis(time.Now().UTC()), channel, event, id, activity)
	}
}
