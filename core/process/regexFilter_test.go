package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {

	t.Run("forwards input to next process", func(t *testing.T) {
		mockEndpoint := &MockProcessor{}
		sut := NewRegexFilter(mockEndpoint, "\\d+")
		input := []byte("1234")

		sut.Process(input)

		assert.Equal(t, input, mockEndpoint.Calls[0], "expected input to be forwarded to endpoint")
	})

	t.Run("filters out input", func(t *testing.T) {
		mockEndpoint := &MockProcessor{}
		sut := NewRegexFilter(mockEndpoint, "\\d+")

		input := []byte("abcd")

		sut.Process(input)

		assert.Equal(t, 0, len(mockEndpoint.Calls), "expected input have been filtered and not forwarded to endpoint")
	})
}

type MockProcessor struct {
	Calls [][]byte
}

func (mock *MockProcessor) Process(input []byte) {
	mock.Calls = append(mock.Calls, input)
}
