package host

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
)

func Ping(uri any) *core.Status {
	return messaging.Ping(Exchange, uri)
}
