package test

import (
	"github.com/advanced-go/common/messaging"
)

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test", messaging.NewEmissaryChannel(true))
	DefaultTracer.Trace(nil, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	DefaultTracer.Trace(a, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	//Output:
	//OnTrace() -> <nil> : event:shutdown agent shutdown
	//OnTrace() -> agent:test : event:shutdown agent shutdown

}
