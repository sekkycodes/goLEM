package discover

import (
	"log"
	"net"
	"strconv"
	"time"
)

type Listener interface {
	Listen(started chan bool, addresses chan string) (err error)
	Stop()
}

type Discoverer struct {
	Port int
	stop chan bool
}

func (d *Discoverer) Listen(started chan bool, addresses chan string) (err error) {
	log.Println("Discoverer starting...")
	portStr := strconv.Itoa(d.Port)
	pc, err := net.ListenPacket("udp", ":"+portStr)
	if err != nil {
		return err
	}
	defer pc.Close()
	started <- true

	buffer := make([]byte, 1024)
	for {
		select {
		case <-d.stop:
			return nil
		default:
			length, addr, err := pc.ReadFrom(buffer)
			if err != nil {
				return err
			}

			deadline := time.Now().Add(10 * time.Millisecond)
			err = pc.SetWriteDeadline(deadline)
			if err != nil {
				return err
			}

			msg := string(buffer[:length])
			log.Printf("received message on discover port %v: %v", d.Port, msg)

			_, e := pc.WriteTo([]byte("ok"), addr)
			if e != nil {
				return err
			}

			addresses <- addr.String()
		}
	}
}
