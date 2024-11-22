package messaging

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
)

func EventErrorStatus(agentId string, msg *Message) *core.Status {
	err := errors.New(fmt.Sprintf("error: message event:%v is invalid for agent:%v", msg.Event(), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}
