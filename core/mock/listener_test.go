package mock

import (
	"testing"

	"net"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	sut := buildSut()

	go func() {
		sut.Start()
	}()

	conn, err := net.Dial("tcp", ":77551")
	if err == nil {
		t.Fatal(err)
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
		Port:        77551,
		ConnType:    "tcp",
		Invocations: inv,
	}
}
