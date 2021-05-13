// Package mock provides mock classes for golems, as well as messaging and communication between golems
// It is intended to be used for integration and higher-level testing (i.e. communication flow between golems)
package mock

import (
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

// Mocked Listener will record any messages sent to a specific port for later verification
type MockedListener struct {
	Port        int
	ConnType    string
	Invocations []string
}

// Start opens the port and records messages sent.
// To stop listening and free the port send "stop" to the port
func (l *MockedListener) Start(started chan bool, stopReceived chan bool) (err error) {
	log.Printf("starting mock listener on port %v for %v connections", l.Port, l.ConnType)
	portStr := strconv.Itoa(l.Port)
	listen, err := net.Listen(l.ConnType, ":"+portStr)
	if err != nil {
		return err
	}
	defer listen.Close()
	started <- true

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			return err
		}

		totalStr := string(buf[:])
		log.Printf("received message %v on port %v", totalStr, l.Port)

		for _, msg := range strings.Split(totalStr, "#") {
			if msg == "stop" {
				log.Printf("stopping listening on port %v", l.Port)
				stopReceived <- true
				return nil
			}

			log.Printf("appending '%v'", msg)
			l.Invocations = append(l.Invocations, msg)
		}
	}
}
