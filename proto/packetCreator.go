package proto

func MakeACKPacket() *ACKPacket {
	return &ACKPacket{
		PacketHead:  PacketHead,
		PacketTrail: PacketTrail,
	}
}
