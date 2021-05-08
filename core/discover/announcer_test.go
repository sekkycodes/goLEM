package discover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeNet struct {
}

type PacketConnMock struct {
	Invocations []string
}

func TestAnnounce(t *testing.T) {
	sut := buildSut()
	message := []byte("dummy message")

	err := sut.Announce(message)

	assert.Nil(t, err)
}

func buildSut() Announcer {
	return Announcer{
		BroadcastIp: "192.168.2.255",
		Port:        12345,
	}
}
