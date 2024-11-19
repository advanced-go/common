package test

import (
	"fmt"
	"github.com/advanced-go/common/messaging"
)

var (
	DefaultTracer = new(defaultTracer)
)

type defaultTracer struct{}

func (d *defaultTracer) Trace(agent messaging.Agent, activity any) {
	fmt.Printf("OnTrace() -> %v : %v\n", agent.Uri(), activity)
}
