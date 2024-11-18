package messaging

import "github.com/advanced-go/common/core"

type Error interface {
	Exception(status *core.Status) *core.Status
}
