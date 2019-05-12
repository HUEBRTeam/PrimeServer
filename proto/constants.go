package proto

const (
	PacketHead  = 0x00000001
	PacketTrail = 0x0FFFFFFF
)

// Packet Types
const (
	PacketGameOver           = 0x01000001
	PacketACK                = 0x01000002
	PacketLogin              = 0x01000003
	PacketProfile            = 0x01000004
	PacketProfileBusy        = 0x01000005
	PacketRequestWorldBest   = 0x01000008
	PacketWorldBest          = 0x01000009
	PacketRequestRankMode    = 0x0100000A
	PacketRankMode           = 0x0100000B
	PacketRequestLevelUpInfo = 0x0100000C
	PacketLevelUpInfo        = 0x0100000D
	PacketScoreBoard         = 0x0100000E
	PacketEnterProfile       = 0x0100000F
	PacketBye                = 0x01000010
	PacketMachineInfo        = 0x01000011
)
