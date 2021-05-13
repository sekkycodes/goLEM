// Discover other golems in the network, announce presence to other golems
package discover

import (
	"net"
	"strconv"
)

// Announces presence to other golems in the network.
// Those golems are expected to acknowledge received information.
// The acknowledgement is handled asynchronously.
type Announcer struct {
	BroadcastIp string
	Port        int
}

// Announce broadcasts into the network to other golems
func (announcer *Announcer) Announce(message []byte) (err error) {
	portStr := ":" + strconv.Itoa(announcer.Port)
	pc, err := net.ListenPacket("udp4", portStr)
	if err != nil {
		return err
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp4", announcer.BroadcastIp+portStr)
	if err != nil {
		return err
	}

	_, err = pc.WriteTo(message, addr)
	return err
}
