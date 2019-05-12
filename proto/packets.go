package proto

import "unsafe"

type ACKPacket struct {
	PacketHead  uint32 // 0x0000001
	PacketType  uint32 // 0x04 0x1000002 // 0x3000000
	PacketTrail uint32 // 0x08 0xFFFFFFF // 0xFFFF
}

type MachineInfoPacket struct {
	PacketHead     uint32        //    0x00 0x0000001
	PacketType     uint32        //    0x04 0x1000011
	MachineID      uint32        //    0x08
	DongleID       uint32        //    0x0C
	CountryID      uint32        //    0x10
	MacAddress     PIUMacAddress //    0x14
	Version        PIUString12   //    0x28
	Processor      PIUString128  //    0x34
	MotherBoard    PIUString128  //    0xb4
	GraphicsCard   PIUString128  //    0x134
	HDDSerial      PIUString32   //    0x1b4
	USBMode        PIUString128  //    Mode 1.0 / Mode 1.1 / Mode 2.0
	Memory         uint32        //    0x254
	ConfigMagic    uint32        //    0x258
	Unk3           uint32        //    0x25c
	Unk4           uint32        //    0x260
	Unk5           uint32        //    0x264
	Unk6           uint32        //    0x268
	Unk7           uint32        //    0x26c
	Unk8           uint32        //    0x270
	Unk9           uint32        //    0x274
	Unk10          uint32        //    0x278
	Unk11          uint32        //    0x27c
	Unk12          uint32
	Unk13          uint32
	Unk14          uint32
	Unk15          uint32
	Unk16          uint32
	Unk17          uint32
	Unk18          uint32
	Unk19          uint32
	Unk20          [76]uint8
	NetworkAddress PIUString16
}

type ScoreBoardPacket struct {
	PacketHead  uint32      //    0x00 0x0000001
	PacketType  uint32      //    0x04 0x100000E
	SongID      uint32      //    0x08
	ChartLevel  uint16      //    0x0C
	Type        uint8       //    0x0E
	Flag        uint8       //    0x0F
	Score       uint32      //    0x10
	RealScore0  uint32      //    0x14
	Unk0        [16]uint8   //    0x18
	RealScore1  uint32      //    Same as SongScore0, dafuq?
	Grade       uint32      //    0x2C
	Kcal        float32     //    0x30
	Perfect     uint32      //    0x34
	Great       uint32      //    0x38
	Good        uint32      //    0x3c
	Bad         uint32      //    0x40
	Miss        uint32      //    0x44
	MaxCombo    uint32      //    0x48
	EXP         uint16      //    0x4c
	PP          uint16      //    0x4e
	Unk1        uint16      //    0x50
	Unk2        uint16      //    0x52
	Unk3        uint32      //    0x54
	Unk4        uint32      //    0x58
	Unk5        uint32      //    0x5c
	Unk6        uint32      //    0x60
	GameVersion PIUString12 //    0x64
	trailing0   uint32      //    0xFFFFFF
	trailing1   uint32      //   0xB21
}

type LoginPacket struct {
	PacketHead uint32      //    0x00 0x0000001
	PacketType uint32      //    0x04 0x1000003
	Unk0       uint32      //    0x08
	Unk1       uint32      //    0x0C
	AccessCode PIUString32 //  Hex String
	Unk2       uint32
}

type ProfilePacket struct {
	PacketHead  uint32       //    0x00 0x0000001
	PacketType  uint32       //    0x04 0x1000004
	Unk0        uint32       //    0x08
	AccessCode  PIUString32  //    0x0C
	Unk1        uint32       //    0x10
	Nickname    PIUNickname  //    0x30
	Unk2        uint32       //    0x3C
	Unk3        uint16       //    0x40
	Level       uint16       //    0x42
	EXP         uint32       //    0x44
	Unk4        uint32       //    0x48
	PP          uint32       //    0x4C
	Unk5        [20]uint8    //    0x50
	RunningStep uint32       //    0x64
	Unk6        [32]uint8    //    0x68
	Scores      [4384]UScore //    0x88
}

type UScore struct {
	SongID     uint32 //    0x00
	ChartLevel uint8  //    0x04
	Unk0       uint8  //    0x05
	Unk1       uint16 //    0x06
	Score      uint32 //    0x08
	RealScore  uint32 //   Maybe
	Unk2       uint32 //    0x10
}

type RequestLevelUpInfoPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x100000C
	Unk0       uint32 //    0x08 0xBD5
}

type LevelUpInfoPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x100000D
	Unk0       uint32 //    0x08 0xBD5
	Level      uint32 //    0x0C
}

type GameOverPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x1000001
	Unk0       uint32 //    0x08 0xBD5
}

type WorldBestPacket struct {
	PacketHead  uint32 //    0x00 0x00000002
	PacketType  uint32 //    0x04 0x10000009
	Unk0        uint32 //    0x08 5056
	Unk1        uint32 //    0x0C 0x0000000F
	Unk2        uint32 //    0x10 674200
	Unk3        uint32 //    0x14 0x00000000
	Unk4        uint32 //    0x18 0x00000000
	WorldScores [4095]WorldBestScore
	Unk5        uint32 //    0x?? 0x00000000
	Unk6        uint32 //    0x?? 0x00000000
	PacketTrail uint32 //    0x?? 0x00000000
}

type WorldBestScore struct {
	SongID     uint32      //
	ChartLevel uint16      //
	ChartMode  uint16      //
	Score      uint32      //
	Unk0       uint32      //
	Unk1       uint32      //
	Nickname   PIUNickname //
}

const ACKPacketLength = int(unsafe.Sizeof(ACKPacket{}))
const MachineInfoPacketLength = int(unsafe.Sizeof(MachineInfoPacket{}))
const MachineInfoPacketLength = int(unsafe.Sizeof(MachineInfoPacket{}))
const ScoreBoardPacketLength = int(unsafe.Sizeof(ScoreBoardPacket{}))
const LoginPacketLength = int(unsafe.Sizeof(LoginPacket{}))
const ProfilePacketLength = int(unsafe.Sizeof(ProfilePacket{}))
const UScoreLength = int(unsafe.Sizeof(UScore{}))
const RequestLevelUpInfoPacketLength = int(unsafe.Sizeof(RequestLevelUpInfoPacket{}))
const LevelUpInfoPacketLength = int(unsafe.Sizeof(LevelUpInfoPacket{}))
const GameOverPacketLength = int(unsafe.Sizeof(GameOverPacket{}))
const WorldBestPacketLength = int(unsafe.Sizeof(WorldBestPacket{}))
const WorldBestScoreLength = int(unsafe.Sizeof(WorldBestScore{}))

var BiggestPacket = 0

func init() {
	packetLens := []int{
		ACKPacketLength, MachineInfoPacketLength, MachineInfoPacketLength,
		ScoreBoardPacketLength, LoginPacketLength, ProfilePacketLength, RequestLevelUpInfoPacketLength,
		LevelUpInfoPacketLength, GameOverPacketLength, WorldBestPacketLength,
	}

	for _, v := range packetLens {
		if v > BiggestPacket {
			BiggestPacket = v
		}
	}
}
