package messaging

type Tracer interface {
	Trace(agentId string, content any)
}
