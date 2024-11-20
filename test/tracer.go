package test

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent messaging.Agent, event, activity string) {
	if agent == nil {
		fmt.Printf("test: Trace() -> %v : %v %v\n", agent, event, activity)

	} else {
		fmt.Printf("test: Trace() -> %v : %v %v\n", agent.Uri(), event, activity)
	}
}
