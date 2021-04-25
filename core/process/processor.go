package process

type Processor interface {
	Process(input []byte)
}
