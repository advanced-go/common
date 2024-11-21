package messaging

import (
	"fmt"
)

type Tracer interface {
	Trace(agent Agent, channel, event, activity string)
}

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent Agent, channel, event, activity string) {
	//name := "<nil>"
	//if agent != nil {
	//	name = agent.Uri()
	//}
	if agent == nil {
		fmt.Printf("%v : %v %v %v", agent, channel, event, activity)
	} else {
		fmt.Printf("%v : %v %v %v", agent.Uri(), channel, event, activity)

	}
}
