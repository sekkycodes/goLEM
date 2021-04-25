package consume

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	outputFile = "testdata/output.json"
)

func TestConsume(t *testing.T) {
	sut := buildSut()
	input := []byte("{ \"key\": 1234 }")

	err := sut.Consume(input)

	assert.Nil(t, err)

	os.Remove(outputFile)
}

func buildSut() JsonFile {
	return JsonFile{
		OutputPath: outputFile,
	}
}
