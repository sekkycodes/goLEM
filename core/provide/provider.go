package provide

type Provider interface {
	Provide() (output []byte, err error)
}
