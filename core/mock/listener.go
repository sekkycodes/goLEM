// Package mock provides mock classes for golems, as well as messaging and communication between golems
// It is intended to be used for integration and higher-level testing (i.e. communication flow between golems)
package mock

import (
	"io/ioutil"
	"net"
	"strconv"
)

// Mocked Listener will record any messages sent to a specific port for later verification
type MockedListener struct {
	Port        int
	ConnType    string
	Invocations []string
}

// Start opens the port and records messages sent.
// To stop listening and free the port send "stop" to the port
func (listener *MockedListener) Start() (err error) {
	portStr := ":" + strconv.Itoa(listener.Port)
	l, err := net.Listen(listener.ConnType, portStr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			return err
		}

		msg := string(buf[:])

		if msg == "stop" {
			return nil
		}

		listener.Invocations = append(listener.Invocations, msg)
	}
}
