package test

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
)

type Notifier interface {
	messaging.Notifier
	Status() *core.Status
	Reset()
}

type statusT struct {
	status *core.Status
}

func (s *statusT) Notify(agent any, status *core.Status) *core.Status {
	s.status = status
	return status
}

func (s *statusT) Status() *core.Status {
	return s.status
}

func (s *statusT) Reset() {
	s.status = nil
}

func NewNotifier() Notifier {
	return new(statusT)
}
