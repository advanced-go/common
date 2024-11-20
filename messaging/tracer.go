package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agent, activity any)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent, activity any) {
	name := "<nil>"
	a := AgentCast(agent)
	if a != nil {
		name = a.Uri()
	}
	fmt.Printf("%v : %v", name, activity)
}
