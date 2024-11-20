package test

import (
	"fmt"
	"github.com/advanced-go/common/core"
)

func ExampleNewNotifier() {
	n := NewNotifier()

	n.Notify(nil, core.StatusNotFound())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	n.Reset()
	n.Notify(nil, core.StatusNoContent())
	fmt.Printf("test: NewNotifier() -> [status:%v]\n", n.Status())

	//Output:
	//test: NewNotifier() -> [status:Not Found]
	//test: NewNotifier() -> [status:No Content]

}
