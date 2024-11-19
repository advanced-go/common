package messaging

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"net/http"
	"reflect"
)

var (
	TestDispatch = new(testDispatch)
)

type testDispatch struct{}

func (t *testDispatch) OnTick(agent any, src *Ticker) {
	fmt.Printf("OnTick() -> %v : %v", DispatchName(agent), DispatchName(src))
}

func (t *testDispatch) OnMessage(agent any, msg *Message, src *Channel) {
	fmt.Printf("OnMsg() -> %v : %v %v", DispatchName(agent), DispatchName(msg), DispatchName(src))
}
func (t *testDispatch) OnError(agent any, status *core.Status) *core.Status {
	fmt.Printf("OnError() -> %v : %v", DispatchName(agent), DispatchName(status))
	return status
}

func DispatchName(t any) string {
	if t == nil {
		return "<nil>"
	}
	switch ptr := t.(type) {
	case Agent:
		return ptr.Uri()
	case *Ticker:
		return ptr.Name()
	case *Channel:
		return ptr.Name()
	case *Message:
		return ptr.Event()
	case *core.Status:
		return ptr.String()
	default:
		return fmt.Sprintf("%v", reflect.TypeOf(t))
	}
}

func ExampleDispatchName() {
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(nil))

	a := NewAgent("agent-test", nil)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(a))

	t := NewTicker("ticker-test", 100)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(t))

	c := NewChannel("channel-test", false)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(c))

	m := NewControlMessage("", "", "event-test")
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(m))

	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(core.StatusNotFound()))

	r := new(http.Response)
	fmt.Printf("test: DispatchName() -> %v\n", DispatchName(r))

	//Output:
	//test: DispatchName() -> <nil>
	//test: DispatchName() -> agent-test
	//test: DispatchName() -> ticker-test
	//test: DispatchName() -> channel-test
	//test: DispatchName() -> event-test
	//test: DispatchName() -> Not Found
	//test: DispatchName() -> *http.Response

}

func ExampleMutedDispatcher() {
	status := exampleDispatch(MutedDispatcher)

	fmt.Printf("test: MutedDispatcher() -> [status:%v]\n", status)

	//Output:
	//test: MutedDispatcher() -> [status:OK]

}

func exampleDispatch(d Dispatcher) *core.Status {
	return core.StatusOK()
}
