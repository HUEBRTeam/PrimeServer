package proto

func MakeACKPacket() *ACKPacket {
	return &ACKPacket{
		PacketHead:  PacketHead,
		PacketType:  PacketACK,
		PacketTrail: PacketTrail,
	}
}
