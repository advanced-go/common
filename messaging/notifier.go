package messaging

import "github.com/advanced-go/common/core"

type Notifier interface {
	Notify(agent any, status *core.Status) *core.Status
}

var (
	LogErrorNotifier    = new(logError)
	OutputErrorNotifier = new(outputError)
)

type logError struct{}

func (l *logError) Notify(agent any, status *core.Status) *core.Status {
	var h core.Log
	return h.Handle(status)
}

type outputError struct{}

func (o *outputError) Notify(agent any, status *core.Status) *core.Status {
	var h core.Output
	return h.Handle(status)
}
