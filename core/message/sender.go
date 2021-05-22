package message

// Sender is responsible for transmitting messages reliable to a receiver.
// What receiver the message is sent to is configurable and usually depends on the next step in a messages journey.
type Sender interface {
}
