package test

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent, activity any) {
	name := "<nil>"
	a := messaging.AgentCast(agent)
	if a != nil {
		name = a.Uri()
	}
	fmt.Printf("OnTrace() -> %v : %v\n", name, activity)
}
