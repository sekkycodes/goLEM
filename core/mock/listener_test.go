package mock

import (
	"testing"
	"time"

	"net"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	sut := buildSut()

	go func() {
		sut.Start()
	}()

	// wait 5ms for sut to spin up
	time.Sleep(5000)

	conn, err := net.Dial("tcp", "localhost:17551")
	if err != nil {
		t.Fatal(err)
	}

	if conn == nil {
		t.Fatal("connection is nil")
	}

	defer conn.Close()

	conn.Write([]byte("test"))
	conn.Write([]byte("stop"))

	assert.Equal(t, 1, len(sut.Invocations))
	assert.Equal(t, "test", sut.Invocations[0])
}

func buildSut() MockedListener {
	var inv []string

	return MockedListener{
		Port:        17551,
		ConnType:    "tcp",
		Invocations: inv,
	}
}
