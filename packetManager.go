package PrimeServer

import (
	"github.com/HUEBRTeam/PrimeServer/proto"
	"net"
)

func SendPacket(conn net.Conn, data []byte) {
	conn.Write(proto.EncryptPacket(data))
}

func SendACK(conn net.Conn) {
	// Send SendACK
	ack := proto.MakeACKPacket()
	SendPacket(conn, ack.ToBinary())
}

func SendProfileBusy(conn net.Conn) {
	pb := proto.MakeProfileBusyPacket()
	SendPacket(conn, pb.ToBinary())
}

func SendProfile(conn net.Conn, profile proto.ProfilePacket) {
	SendPacket(conn, profile.ToBinary())
}

func SendLevelUpPacket(conn net.Conn, profileId, level uint32) {
	lu := proto.MakeLevelupInfoPacket(profileId, level)
	SendPacket(conn, lu.ToBinary())
}
