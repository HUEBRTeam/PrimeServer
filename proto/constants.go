package proto

import "fmt"

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
	PacketScoreBoardV2       = 0x01000014
	PacketLoginV2            = 0x01000015
	PacketKeepAlive          = 0x03000000
)

var packetTypeToName = map[uint32]string{
	PacketGameOver:           "Game Over",
	PacketACK:                "ACK",
	PacketLogin:              "Login",
	PacketProfile:            "Profile",
	PacketProfileBusy:        "Profile Busy",
	PacketRequestWorldBest:   "Request World Best",
	PacketWorldBest:          "World Best",
	PacketRequestRankMode:    "Request Rank Mode",
	PacketRankMode:           "Rank Mode",
	PacketRequestLevelUpInfo: "Request Levelup Info",
	PacketLevelUpInfo:        "Levelup Info",
	PacketScoreBoard:         "Score Board",
	PacketEnterProfile:       "Enter Profile",
	PacketBye:                "Bye",
	PacketMachineInfo:        "Machine Info",
	PacketLoginV2:            "Login V2",
	PacketKeepAlive:          "Keep Alive",
	PacketScoreBoardV2:       "ScoreBoard V2",
}

func GetPacketName(t uint32) string {
	if name, ok := packetTypeToName[t]; ok {
		return name
	}

	return fmt.Sprintf("Unknown 0x%x", t)
}
