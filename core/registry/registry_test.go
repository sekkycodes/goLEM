package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersistFile(t *testing.T) {

}

func TestRetrieveAllFiles(t *testing.T) {
	reg := FileReg{
		FilePath: "testdata",
	}

	entries, err := reg.RetrieveAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(entries))
}
