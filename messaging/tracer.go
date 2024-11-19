package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agent Agent, activity any)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent Agent, activity any) {
	fmt.Printf("%v : %v", agent.Uri(), activity)
}
