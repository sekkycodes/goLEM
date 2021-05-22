package discover

import (
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {

	callback := make(chan string)
	started := make(chan bool)
	stop := make(chan bool)

	sut := Discoverer{
		Port: 44821,
		stop: stop,
	}

	go func() {
		err := sut.Listen(started, callback)
		if err != nil {
			assert.Fail(t, "failed to start discoverer")
		}
	}()

	<-started

	go func() {
		runTest(t, stop)
	}()

	address := <-callback

	assert.True(t, strings.HasPrefix(address, "127.0.0.1"))
}

func runTest(t *testing.T, stop chan bool) {
	conn, err := net.Dial("udp", ":44821")
	if err != nil {
		t.Fatal(err)
	}

	if conn == nil {
		t.Fatal("connction is nil")
	}

	defer conn.Close()

	conn.Write([]byte("hi"))
	// give the listener some time to handle message
	time.Sleep(20 * time.Millisecond)
	// stop listener
	stop <- true
}
