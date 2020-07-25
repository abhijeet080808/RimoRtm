// Package udpserver provides a UDPServer
package udpserver

import (
	"fmt"
	"log"
	"net"
)

// UDPServer listens for status messages
type UDPServer struct {
	host string
	port int
	// storing the interface here which holds pointer to the net.UDPConn struct
	packetConn net.PacketConn
}

// New creates a new UDPServer on specified host and port
func New(host string, port int) *UDPServer {
	server := UDPServer{
		host: host,
		port: port,
	}

	log.Println("Starting UDP server at", host, port)
	packetConn, err := net.ListenPacket("udp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalln(err)
	}

	server.packetConn = packetConn

	// Run on a separate goroutine
	go server.ReadPackets()

	return &server
}

// ReadPackets reads all UDP packets are they are received
func (s *UDPServer) ReadPackets() {
	buffer := make([]byte, 1500)

	for {
		// Blocking read call
		n, addr, err := s.packetConn.ReadFrom(buffer)
		if n > 0 {
			log.Println("UDP packet received: bytes:", n, "from:", addr.String())
		}
		if err != nil {
			log.Panicln(err)
		}
	}
}
