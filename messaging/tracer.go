package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agent any, activity any)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent any, activity any) {
	name := "<nil>"
	a := AgentCast(agent)
	if a != nil {
		name = a.Uri()
	}
	fmt.Printf("%v : %v", name, activity)
}
