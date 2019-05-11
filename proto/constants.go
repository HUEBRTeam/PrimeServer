package proto

const (
	PacketHead  = 0x0000001
	PacketTrail = 0xFFFFFFF
)

// Packet Types
const (
	PacketMachineInfo        = 0x1000001
	PacketScoreBoard         = 0x100000E
	PacketLogin              = 0x1000003
	PacketProfile            = 0x1000004
	PacketRequestLevelUpInfo = 0x100000C
	PacketLevelUpInfo        = 0x100000D
	PacketGameOver           = 0x1000001
	PacketWorldBest          = 0x10000009
)
