package provide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvide(t *testing.T) {
	sut := buildSut()
	content, err := sut.Provide("testdata/test.json")

	assert.Nil(t, err, "no error should be set")
	assert.NotNil(t, content, "content should be set")
}

func buildSut() *JsonFileGolem {
	return new(JsonFileGolem)
}
