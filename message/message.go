package message

const (
	UserMessage = iota
	BroadcastMessage
)

type Message struct {
	Type int
	Username string `json:"username"`
	Message string `json:"message"`
}