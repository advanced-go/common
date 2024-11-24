package messaging

import "fmt"

func ExampleTraceDispatch_Trace() {
	fmt.Printf("test: TraceDispatch() -> \n")

	//Output:
	//test: TraceDispatch() ->
}

/*
func ExampleMutedDispatcher() {
	//status := exampleDispatch(MutedDispatcher)

	//fmt.Printf("test: MutedDispatcher() -> [status:%v]\n", status)

	//Output:
	//test: MutedDispatcher() -> [status:OK]

}

func exampleDispatch(d Dispatcher) *core.Status {
	return core.StatusOK()
}


*/
