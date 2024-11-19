package messaging

func ExampleDefaultTracer_Trace() {
	activity := "agent shutdown"
	DefaultTracer.Trace("test: Trace() -> case-officer", activity)

	//Output:
	//test: Trace() -> case-officer : agent shutdown

}
