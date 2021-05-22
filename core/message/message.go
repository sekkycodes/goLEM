// Message package includes all types and methods needed for handling data exchange between golems.
package message

// Messager wraps content and sends it to other golems
type Messager interface {
	// TODO: protobuf message consisting of:
	// - message ID
	// - correlation ID
	// - content ([]byte?)
	// - order number
	// ...
}
