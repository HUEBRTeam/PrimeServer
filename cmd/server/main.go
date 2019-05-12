package main

import (
	"encoding/binary"
	"fmt"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"net"
	"os"
)

const (
	ConnPort = "60010"
	ConnType = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(ConnType, ":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on :" + ConnPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handlePacket(conn net.Conn, packet []byte) {
	dec, ok := proto.DecryptPacket(packet)

	if !ok {
		fmt.Println("Received invalid packet")
		return
	}

	packetType := binary.LittleEndian.Uint32(dec[4:8])

	switch packetType {
	case proto.PacketMachineInfo2:
		fmt.Println("Received Machine Info V2")
	default:
		fmt.Printf("Received packet 0x%x\n", packetType)
	}

	// Send ACK
	ack := proto.MakeACKPacket()
	conn.Write(proto.EncryptPacket(ack.ToBinary()))
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	c := 0
	running := true
	defer conn.Close()
	for running {
		// Read the incoming connection into the buffer.
		n, err := conn.Read(buf[c:])
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			running = false
			return
		}
		c += n

		fmt.Printf("Received %d bytes\n", c)

		if c > 4 {
			plen := int(binary.LittleEndian.Uint32(buf[:4]))

			fmt.Printf("Waiting for %d bytes got %d\n", plen, c)

			if plen > proto.BiggestPacket {
				fmt.Printf("Invalid packet length: %d", plen)
				running = false
				return
			}

			if c >= plen {
				handlePacket(conn, buf[4:plen])
				copy(buf, buf[plen:])
				c -= plen
			}
		}
	}
}
