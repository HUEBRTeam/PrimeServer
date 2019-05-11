package proto

const (
	PacketHead  = 0x0000001
	PacketTrail = 0xFFFFFFF
)

// Packet Types
const (
	PacketMachineInfo        = 0x01000001
	PacketGameOver           = 0x01000001
	PacketACK                = 0x01000002
	PacketLogin              = 0x01000003
	PacketProfile            = 0x01000004
	PacketRequestLevelUpInfo = 0x0100000C
	PacketLevelUpInfo        = 0x0100000D
	PacketScoreBoard         = 0x0100000E
	PacketWorldBest          = 0x10000009
)
