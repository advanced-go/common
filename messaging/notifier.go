package messaging

import "github.com/advanced-go/common/core"

type Notifier interface {
	Notify(status *core.Status) *core.Status
}
