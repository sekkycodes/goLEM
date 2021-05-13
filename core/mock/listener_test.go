package mock

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	sut := buildSut()
	started := make(chan bool)
	stopReceived := make(chan bool)

	go func() {
		sut.Start(started, stopReceived)
	}()

	log.Println("waiting for sut to spin up")
	<-started
	log.Println("sut started")

	go func() {
		runTest(t)
	}()

	<-stopReceived

	log.Printf("messages received: %v", sut.Invocations)

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

func runTest(t *testing.T) {
	conn, err := net.Dial("tcp", ":17551")
	if err != nil {
		t.Fatal(err)
	}

	if conn == nil {
		t.Fatal("connection is nil")
	}

	defer conn.Close()

	conn.Write([]byte("test"))
	conn.Write([]byte("#")) // signal end of message
	conn.Write([]byte("stop"))
	conn.Write([]byte("#")) // signal end of message
}
