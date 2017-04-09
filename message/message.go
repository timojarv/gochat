package message

const (
	Msg = iota
	Broadcast
	Metadata
)

type Message struct {
	Type int `json:"type"`
	Error bool `json:"error"`
	Sender string `json:"sender"`
	Body string `json:"body"`
}

func Meta(body string) Message {
	return Message{
		Type: Metadata,
		Body: body,
	}
}

func Error(body string) Message {
	return Message{
		Type: Metadata,
		Body: body,
		Error: true,
	}
}

func New(sender, body string) Message {
	return Message{
		Type: Msg,
		Body: body,
		Sender: sender,
	}
}