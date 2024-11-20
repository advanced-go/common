package test

func ExampleDefaultTracer_Trace() {
	a := NewAgent("agent:test")
	DefaultTracer.Trace(nil, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	DefaultTracer.Trace(a, "event:shutdown", "agent shutdown")
	//fmt.Printf("\n")

	//Output:
	//test: Trace() -> <nil> : event:shutdown agent shutdown
	//test: Trace() -> agent:test : event:shutdown agent shutdown

}
