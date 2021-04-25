package consume

type Consumer interface {
	Consume(input []byte) (err error)
}
