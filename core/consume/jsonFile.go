package consume

import (
	"io/ioutil"
)

// JsonFile consumes input and writes it to a file path
type JsonFile struct {
	OutputPath string
}

// Consume writes byte array to file
func (consumer *JsonFile) Consume(input []byte) (err error) {
	return ioutil.WriteFile(consumer.OutputPath, input, 0664)
}
