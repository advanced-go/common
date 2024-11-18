package messaging

import (
	"github.com/advanced-go/common/core"
)

const (
	ChannelSize = 16
)

// OnShutdown - add functions to be run on shutdown
type OnShutdown interface {
	Add(func())
}

// Agent - intelligent agent
// TODO : Track agent assignment as part of the URI or separate identifier??
// //Uri() string
//
//	//Message(m *Message)
//	Track agent NID or class/type?
type Agent interface {
	Mailbox
	Finalizer
	Run()
	Shutdown()
}

type OpsAgent interface {
	Agent
	Notifier
	Tracer
	core.ErrorHandler
	//AddActivity(agentId string, content any)
}

func AddShutdown(curr, next func()) func() {
	if next == nil {
		return nil
	}
	if curr == nil {
		curr = next
	} else {
		// !panic
		prev := curr
		curr = func() {
			prev()
			next()
		}
	}
	return curr
}

/*
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in agent.Shutdown() : %v\n", r)
		}
	}()

*/
