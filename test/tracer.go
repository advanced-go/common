package test

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent messaging.Agent, channel, event, activity string) {
	trace(agent, channel, event, activity)
}

func trace(agent messaging.Agent, channel, event, activity string) {
	if agent == nil {
		fmt.Printf("test: Trace() -> %v : [%v] [%v] [%v]\n", agent, channel, event, activity)

	} else {
		fmt.Printf("test: Trace() -> %v : [%v] [%v] [%v]\n", agent.Uri(), channel, event, activity)
	}
}
