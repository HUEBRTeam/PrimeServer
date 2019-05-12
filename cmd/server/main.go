package main

import (
	"encoding/binary"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/quan-to/slog"
	"net"
	"os"
)

const (
	ConnPort = "60010"
	ConnType = "tcp"
)

var log = slog.Scope("PrimeServer")

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(ConnType, ":"+ConnPort)
	if err != nil {
		log.Error("Error listening: %s", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	log.Info("Listening on :%s", ConnPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Error("Error accepting: %s", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func SendPacket(conn net.Conn, data []byte) {
	conn.Write(proto.EncryptPacket(data))
}

func ACK(conn net.Conn) {
	// Send ACK
	ack := proto.MakeACKPacket()
	SendPacket(conn, ack.ToBinary())
}

func handlePacket(conn net.Conn, packet []byte) {
	l := slog.Scope(conn.RemoteAddr().String())
	dec, ok := proto.DecryptPacket(packet)

	if !ok {
		l.Error("Received invalid packet.")
		return
	}

	p, err := proto.DecodePacket(dec)

	if err != nil {
		l.Error("Error parsing packet: %s", err)
		return
	}

	l.Debug("Received Packet: %s", p.GetName())

	switch v := p.(type) {
	case *proto.RequestWorldBestPacket:
		wb := proto.MakeWorldBestPacket(nil)
		SendPacket(conn, wb.ToBinary())
	case *proto.RequestRankModePacket:
		rm := proto.MakeRankModePacket(nil)
		SendPacket(conn, rm.ToBinary())
	case *proto.MachineInfoPacket:
		log.Debug(v)
		ACK(conn)
	default:
		ACK(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	l := slog.Scope(conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	c := 0
	running := true
	defer conn.Close()

	log.Debug("Connection from %s", conn.RemoteAddr().String())

	for running {
		// Read the incoming connection into the buffer.
		n, err := conn.Read(buf[c:])
		if err != nil {
			l.Error("Error reading: %s", err.Error())
			running = false
			return
		}
		c += n

		if c > 4 {
			plen := int(binary.LittleEndian.Uint32(buf[:4]))

			if plen > proto.BiggestPacket {
				l.Error("Invalid packet length: %d", plen)
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
