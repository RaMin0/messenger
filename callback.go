package messenger

type callbacks struct {
	Message  messageCallbackFunc
	Postback postbackCallbackFunc
	Location locationCallbackFunc
}

type base struct {
	Sender *User
}

// Message type
type Message struct {
	base
	MID         string
	Seq         float64
	Text        string
	Attachments []*struct {
		Title   string
		URL     string
		Type    string
		Payload *struct {
			Coordinates *Location
		}
	}
}

// HasAttachments func
func (m *Message) HasAttachments() bool {
	return m.Attachments != nil && len(m.Attachments) > 0
}

// IsLocation func
func (m *Message) IsLocation() bool {
	if !m.HasAttachments() {
		return false
	}

	for _, a := range m.Attachments {
		if a.Payload != nil && a.Payload.Coordinates != nil {
			return true
		}
	}
	return false
}

type messageCallbackFunc func(*Messenger, *Message)

// HandleMessage func
func (ms *Messenger) HandleMessage(fn messageCallbackFunc) *Messenger {
	ms.callbacks.Message = fn
	return ms
}
func (ms *Messenger) handleMessage(m *Message) {
	if ms.callbacks.Message != nil {
		ms.callbacks.Message(ms, m)
	}
}

// Location type
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type locationCallbackFunc func(*Messenger, *Location)

// HandleLocation func
func (ms *Messenger) HandleLocation(fn locationCallbackFunc) *Messenger {
	ms.callbacks.Location = fn
	return ms
}
func (ms *Messenger) handleLocation(l *Location) {
	if ms.callbacks.Location != nil {
		ms.callbacks.Location(ms, l)
	}
}

// Postback type
type Postback struct {
	base
	Payload string
}

type postbackCallbackFunc func(*Messenger, *Postback)

// HandlePostback func
func (ms *Messenger) HandlePostback(fn postbackCallbackFunc) *Messenger {
	ms.callbacks.Postback = fn
	return ms
}
func (ms *Messenger) handlePostback(p *Postback) {
	if ms.callbacks.Postback != nil {
		ms.callbacks.Postback(ms, p)
	}
}

func (ms *Messenger) handle(_ *entry, g *Messaging) {
	if g.Message != nil {
		m := g.Message
		m.Sender = g.Sender
		ms.handleMessage(m)
	}

	if g.Postback != nil {
		p := g.Postback
		p.Sender = g.Sender
		ms.handlePostback(p)
	}
}
