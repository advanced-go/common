package test

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"reflect"
)

var (
	Dispatcher = new(testDispatch)
)

type testDispatch struct{}

func (t *testDispatch) OnTick(agent any, src *messaging.Ticker) {
	fmt.Printf("OnTick() -> %v : %v", DispatchName(agent), DispatchName(src))
}

func (t *testDispatch) OnMessage(agent any, msg *messaging.Message, src *messaging.Channel) {
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
	case messaging.Agent:
		return ptr.Uri()
	case *messaging.Ticker:
		return ptr.Name()
	case *messaging.Channel:
		return ptr.Name()
	case *messaging.Message:
		return ptr.Event()
	case *core.Status:
		return ptr.String()
	default:
		return fmt.Sprintf("%v", reflect.TypeOf(t))
	}
}
