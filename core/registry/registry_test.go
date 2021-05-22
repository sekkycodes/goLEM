package registry

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPersistFile(t *testing.T) {
	regDir := t.TempDir()
	reg := FileReg{
		FilePath: regDir,
	}

	id := uuid.MustParse("81484e36-07a2-49b8-b572-4db311c93fd8")
	reg.Persist(RegEntry{
		id:         id,
		uri:        "127.0.0.1",
		discovered: 1621675518341,
	})

	entry, err := reg.Retrieve(id)

	assert.Nil(t, err)
	assert.Equal(t, id, entry.id)
}

func TestRetrieveAllFiles(t *testing.T) {
	reg := FileReg{
		FilePath: "testdata",
	}

	entries, err := reg.RetrieveAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(entries))
}
