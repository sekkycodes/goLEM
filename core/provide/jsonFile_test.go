package provide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvide(t *testing.T) {
	sut := buildSut()
	content, err := sut.Provide()

	assert := assert.New(t)
	assert.Nil(err, "no error should be set")
	assert.NotNil(content, "content should be set")
}

func buildSut() JsonFile {
	return JsonFile{
		Source: "testdata/test.json",
	}
}
