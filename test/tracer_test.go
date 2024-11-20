package test

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test")
	DefaultTracer.Trace(nil, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	DefaultTracer.Trace(a, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	//Output:
	//OnTrace() -> <nil> : event:shutdown agent shutdown
	//OnTrace() -> agent:test : event:shutdown agent shutdown

}
