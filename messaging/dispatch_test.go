package messaging

import (
	"fmt"
	"github.com/advanced-go/common/core"
)

func exampleDispatch(d Dispatcher) *core.Status {
	return core.StatusOK()
}

func ExampleMutedDispatcher() {
	status := exampleDispatch(MutedDispatcher)

	fmt.Printf("test: MutedDispatcher() -> [status:%v]\n", status)

	//Output:
	//test: MutedDispatcher() -> [status:OK]

}
