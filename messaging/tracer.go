package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agentId string, activity any)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agentId string, activity any) {
	fmt.Printf("%v : %v", agentId, activity)
}
