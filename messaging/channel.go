package messaging

const (
	EmissaryChannel = "EMISSARY"
	MasterChannel   = "MASTER"
)

type Channel struct {
	enabled bool
	name    string
	C       chan *Message
}

func NewChannel(name string, enabled bool) *Channel {
	c := new(Channel)
	c.name = name
	c.enabled = enabled
	c.C = make(chan *Message, ChannelSize)
	return c
}

func NewEmissaryChannel(enabled bool) *Channel {
	return NewChannel(EmissaryChannel, enabled)
}

func NewMasterChannel(enabled bool) *Channel {
	return NewChannel(MasterChannel, enabled)
}

func (c *Channel) String() string  { return c.Name() }
func (c *Channel) Name() string    { return c.name }
func (c *Channel) IsEnabled() bool { return c.enabled }
func (c *Channel) Enable()         { c.enabled = true }
func (c *Channel) Disable()        { c.enabled = false }

func (c *Channel) Close() {
	if c.C != nil {
		close(c.C)
		c.C = nil
	}
}

func (c *Channel) Send(m *Message) {
	if m != nil && c.enabled {
		c.C <- m
	}
}
