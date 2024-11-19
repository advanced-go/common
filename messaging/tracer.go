package messaging

type Tracer interface {
	Trace(agentId string, activity any)
}
