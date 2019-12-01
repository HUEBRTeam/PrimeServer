package proto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type SimplePacket struct {
	PacketHead uint32 `json:"-"`
	PacketType uint32 `json:"-"`
}

type ACKPacket struct {
	SimplePacket
	PacketTrail uint32 `json:"-"`
}

type ProfileBusyPacket struct {
	SimplePacket
	Unk0 uint32
}
type ByePacket struct {
	SimplePacket
	ProfileID uint32
}

type EnterProfilePacket struct {
	SimplePacket
	PlayerID  uint32
	MachineID uint32
	ProfileID uint32
}

type RequestWorldBestPacket struct {
	SimplePacket
}

type RequestRankModePacket struct {
	SimplePacket
}

type KeepAlivePacket struct {
	SimplePacket
	PacketTrail uint32 `json:"-"`
}

type UnknownPacket0 struct {
	PacketHead uint32 `json:"-"` // 0x0D
	PacketType uint32 `json:"-"` // 0x100001C
	Trailing   uint32 `json:"-"`
}

type UnknownPacket1 struct {
	SimplePacket
}

type MachineInfoPacket struct {
	PacketHead     uint32        `json:"PacketHead"`//    0x00 0x0000001
	PacketType     uint32        `json:"PacketType"`//    0x04 0x1000011
	MachineID      uint32        `json:"MachineID"`//    0x08
	DongleID       uint32        `json:"DongleID"`//    0x0C
	CountryID      uint32        `json:"CountryID"`//    0x10
	MacAddress     PIUMacAddress `json:"MacAddress"`//    0x14
	Version        PIUString12   `json:"Version"`//    0x28
	Processor      PIUString128  `json:"Processor"`//    0x34
	MotherBoard    PIUString128  `json:"MotherBoard"`//    0xb4
	GraphicsCard   PIUString128  `json:"GraphicsCard"`//    0x134
	HDDSerial      PIUString32   `json:"HDDSerial"`//    0x1b4
	USBMode        PIUString128  `json:"USBMode"`//    Mode 1.0 / Mode 1.1 / Mode 2.0
	Memory         uint32        `json:"Memory"`//    0x254
	ConfigMagic    uint32        `json:"ConfigMagic"`//    0x258
	Unk3           uint32        `json:"-"` //    0x25c
	Unk4           uint32        `json:"-"` //    0x260
	Unk5           uint32        `json:"-"` //    0x264
	Unk6           uint32        `json:"-"` //    0x268
	Unk7           uint32        `json:"-"` //    0x26c
	Unk8           uint32        `json:"-"` //    0x270
	Unk9           uint32        `json:"-"` //    0x274
	Unk10          uint32        `json:"-"` //    0x278
	Unk11          uint32        `json:"-"` //    0x27c
	Unk12          uint32        `json:"-"`
	Unk13          uint32        `json:"-"`
	Unk14          uint32        `json:"-"`
	Unk15          uint32        `json:"-"`
	Unk16          uint32        `json:"-"`
	Unk17          uint32        `json:"-"`
	Unk18          uint32        `json:"-"`
	Unk19          [76]uint8     `json:"-"`
	NetworkAddress PIUString16   `json:"NetworkAddress"`
}

func (p *MachineInfoPacket) String() string {
	s := "Machine Packet\n"

	s += fmt.Sprintf("\tPacketHead: %d (0x%x)\n", p.PacketHead, p.PacketHead)
	s += fmt.Sprintf("\tPacketType: %d (0x%x)\n", p.PacketType, p.PacketType)
	s += fmt.Sprintf("\tMachineID: %d (0x%x)\n", p.MachineID, p.MachineID)
	s += fmt.Sprintf("\tDongle ID: %d (0x%x)\n", p.DongleID, p.DongleID)
	s += fmt.Sprintf("\tCountry ID: %d (0x%x)\n", p.CountryID, p.CountryID)
	s += fmt.Sprintf("\tMac Address: %s\n", p.MacAddress)
	s += fmt.Sprintf("\tVersion: %s\n", p.Version)
	s += fmt.Sprintf("\tProcessor: %s\n", p.Processor)
	s += fmt.Sprintf("\tMother Board: %s\n", p.MotherBoard)
	s += fmt.Sprintf("\tGraphics Card: %s\n", p.GraphicsCard)
	s += fmt.Sprintf("\tHDD Serial: %s\n", p.HDDSerial)
	s += fmt.Sprintf("\tUSB Mode: %s\n", p.USBMode)
	s += fmt.Sprintf("\tMemory: %d\n", p.Memory)
	s += fmt.Sprintf("\tConfig Magic: %d (0x%x)\n", p.ConfigMagic, p.ConfigMagic)
	s += fmt.Sprintf("\tNet Address: %s\n", p.NetworkAddress)

	s += fmt.Sprintf("\tUnknown uint32_t  3: %d (0x%x)\n", p.Unk3, p.Unk3)
	s += fmt.Sprintf("\tUnknown uint32_t  4: %d (0x%x)\n", p.Unk4, p.Unk4)
	s += fmt.Sprintf("\tUnknown uint32_t  5: %d (0x%x)\n", p.Unk5, p.Unk5)
	s += fmt.Sprintf("\tUnknown uint32_t  6: %d (0x%x)\n", p.Unk6, p.Unk6)
	s += fmt.Sprintf("\tUnknown uint32_t  7: %d (0x%x)\n", p.Unk7, p.Unk7)
	s += fmt.Sprintf("\tUnknown uint32_t  8: %d (0x%x)\n", p.Unk8, p.Unk8)
	s += fmt.Sprintf("\tUnknown uint32_t  9: %d (0x%x)\n", p.Unk9, p.Unk9)
	s += fmt.Sprintf("\tUnknown uint32_t 10: %d (0x%x)\n", p.Unk10, p.Unk10)
	s += fmt.Sprintf("\tUnknown uint32_t 11: %d (0x%x)\n", p.Unk11, p.Unk11)
	s += fmt.Sprintf("\tUnknown uint32_t 12: %d (0x%x)\n", p.Unk12, p.Unk12)
	s += fmt.Sprintf("\tUnknown uint32_t 13: %d (0x%x)\n", p.Unk13, p.Unk13)
	s += fmt.Sprintf("\tUnknown uint32_t 14: %d (0x%x)\n", p.Unk14, p.Unk14)
	s += fmt.Sprintf("\tUnknown uint32_t 15: %d (0x%x)\n", p.Unk15, p.Unk15)
	s += fmt.Sprintf("\tUnknown uint32_t 16: %d (0x%x)\n", p.Unk16, p.Unk16)
	s += fmt.Sprintf("\tUnknown uint32_t 17: %d (0x%x)\n", p.Unk17, p.Unk17)
	s += fmt.Sprintf("\tUnknown uint32_t 18: %d (0x%x)\n", p.Unk18, p.Unk18)
	s += fmt.Sprintf("\tUnknown String: %s\n", hex.EncodeToString(p.Unk19[:]))

	return s
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
	Unk0        [16]uint8   `json:"-"` //    0x18
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
	RunningStep uint16      //    0x50
	Unk2        uint16      `json:"-"` //    0x52
	Unk3        uint32      `json:"-"` //    0x54
	Unk4        uint32      `json:"-"` //    0x58
	Unk5        uint32      `json:"-"` //    0x5c
	RushSpeed   float32     //    0x60
	GameVersion PIUString12 //    0x64
	MachineID   uint32      //    0xFFFFFF
	ProfileID   uint32      //   0xB21
}

type ScoreBoardPacket2 struct {
	PacketHead    uint32      `json:"PacketHead"`//    0x00 0x0000001
	PacketType    uint32      `json:"PacketType"`//    0x04 0x1000014
	SongID        uint32      `json:"SongID"`//    0x08
	ChartLevel    uint16      `json:"ChartLevel"`//    0x0C
	Type          uint8       `json:"Type"`//    0x0E
	Flag          uint8       `json:"Flag"`//    0x0F
	Score         uint32      `json:"Score"`//    0x10
	RealScore0    uint32      `json:"RealScore0"`//    0x14
	Unk0          [16]uint8   `json:"-"` //    0x18
	RealScore1    uint32      `json:"RealScore1"`//    Same as SongScore0, dafuq?
	Grade         uint32      `json:"Grade"`//    0x2C
	Kcal          float32     `json:"Kcal"`//    0x30
	Perfect       uint32      `json:"Perfect"`//    0x34
	Great         uint32      `json:"Great"`//    0x38
	Good          uint32      `json:"Good"`//    0x3c
	Bad           uint32      `json:"Bad"`//    0x40
	Miss          uint32      `json:"Miss"`//    0x44
	MaxCombo      uint32      `json:"MaxCombo"`//    0x48
	EXP           uint16      `json:"EXP"`//    0x4c
	PP            uint16      `json:"PP"`//    0x4e
	RunningStep   uint16      `json:"RunningStep"`//    0x50
	Unk2          uint16      `json:"-"` //    0x52
	Modifiers     uint32      `json:"Modifiers"`//    0x54
	Unk4          uint32      `json:"-"` //    0x58
	NoteSkinSpeed uint32      `json:"NoteSkinSpeed"`//    0x5c // Contains scroll speed, 0x14 for 5x and 0x0C for 3x
	RushSpeed     float32     `json:"RushSpeed"`//    0x60
	GameVersion   PIUString12 `json:"GameVersion"`//    0x64
	MachineID     uint32      `json:"MachineID"`//    0xFFFFFF
	ProfileID     uint32      `json:"ProfileID"`//   0xB21
	SongCategory  uint32      `json:"SongCategory"`
	Unk7          uint32 `json:"-"`
}

func (p *ScoreBoardPacket2) String() string {
	s := "ScoreBoard V2: \n"

	s += fmt.Sprintf("\tSongID: %d (%x)\n", p.SongID, p.SongID)
	s += fmt.Sprintf("\tChart Level: %d (%x)\n", p.ChartLevel, p.ChartLevel)
	s += fmt.Sprintf("\tType: %d (%x)\n", p.Type, p.Type)
	s += fmt.Sprintf("\tFlag: %d (%x)\n", p.Flag, p.Flag)
	s += fmt.Sprintf("\tScore: %d\n", p.Score)
	s += fmt.Sprintf("\tRealScore0: %d\n", p.RealScore0)
	s += fmt.Sprintf("\tUnk0: %d (%x)\n", p.Unk0, p.Unk0)
	s += fmt.Sprintf("\tRealScore1: %d\n", p.RealScore1)
	s += fmt.Sprintf("\tGrade: %d (%x)\n", p.Grade, p.Grade)
	s += fmt.Sprintf("\tKcal: %f\n", p.Kcal)
	s += fmt.Sprintf("\tPerfect: %d\n", p.Perfect)
	s += fmt.Sprintf("\tGreat: %d\n", p.Great)
	s += fmt.Sprintf("\tGood: %d\n", p.Good)
	s += fmt.Sprintf("\tBad: %d\n", p.Bad)
	s += fmt.Sprintf("\tMiss: %d\n", p.Miss)
	s += fmt.Sprintf("\tMaxCombo: %d\n", p.MaxCombo)
	s += fmt.Sprintf("\tEXP: %d\n", p.EXP)
	s += fmt.Sprintf("\tPP: %d\n", p.PP)
	s += fmt.Sprintf("\tRunningStep: %d\n", p.RunningStep)
	s += fmt.Sprintf("\tUnk2: %d (%x)\n", p.Unk2, p.Unk2)
	s += fmt.Sprintf("\tModifiers: %d (%x)\n", p.Modifiers, p.Modifiers)
	s += fmt.Sprintf("\tUnk4: %d (%x)\n", p.Unk4, p.Unk4)
	s += fmt.Sprintf("\tNoteSkinSpeed: %d (%x)\n", p.NoteSkinSpeed, p.NoteSkinSpeed)
	s += fmt.Sprintf("\tRushSpeed: %f\n", p.RushSpeed)
	s += fmt.Sprintf("\tGameVersion: %s\n", p.GameVersion.String())
	s += fmt.Sprintf("\tMachineID: %d (%x)\n", p.MachineID, p.MachineID)
	s += fmt.Sprintf("\tProfileID: %d (%x)\n", p.ProfileID, p.ProfileID)
	s += fmt.Sprintf("\tSongCategory: %d (%x)\n", p.SongCategory, p.SongCategory)
	s += fmt.Sprintf("\tUnk7: %d (%x)\n", p.Unk7, p.Unk7)

	return s
}

type LoginPacket struct {
	PacketHead  uint32      //    0x00 0x0000001
	PacketType  uint32      //    0x04 0x1000003
	PlayerID    uint32      //    0x08
	MachineID   uint32      //    0x0C
	AccessCode  PIUString32 //  Hex String
	PacketTrail uint32
}

type LoginPacketV2 struct {
	PacketHead  uint32      //    0x00 0x0000004
	PacketType  uint32      //    0x04 0x1000015
	PlayerID    uint32      //    0x08
	MachineID   uint32      //    0x0C
	AccessCode  PIUString32 //  Hex String
	Unk0        uint32      `json:"-"`
	GameVersion PIUString12
	PacketTrail uint32
}

type ProfilePacket struct {
	PacketHead    uint32      `json:"-"` //    0x00 0x0000001
	PacketType    uint32      `json:"-"` //    0x04 0x1000004
	PlayerID      uint32      `json:"PlayerID"`//    0x08
	AccessCode    PIUString32 `json:"AccessCode"`//    0x0C
	Unk0          uint32      `json:"-"`
	Nickname      PIUNickname `json:"Nickname"`//    0x30
	ProfileID     uint32      `json:"ProfileID"`//    0x10
	CountryID     uint8       `json:"CountryID"`//    0x3C
	Avatar        uint8       `json:"Avatar"`//    0x40
	Level         uint8       `json:"Level"`//    0x42
	Unk1          uint8       `json:"-"`
	EXP           uint64      `json:"EXP"`
	PP            uint64      `json:"PP"`
	RankSingle    uint64      `json:"RankSingle"`
	RankDouble    uint64      `json:"RankDouble"`
	RunningStep   uint64      `json:"RunningStep"`
	PlayCount     uint32      `json:"PlayCount"`
	Kcal          float32     `json:"Kcal"`
	Modifiers     uint64      `json:"Modifiers"`
	NoteSkinSpeed uint32      `json:"NoteSkinSpeed"`
	RushSpeed     float32     `json:"RushSpeed"`
	Unk2          uint32       `json:"-"`
	Scores        [4384]UScore `json:"Scores"`//    0x88
}

type UScore struct {
	SongID       uint32 `json:"SongID"`//    0x00
	ChartLevel   uint8  `json:"ChartLevel"`//    0x04
	Unk0         uint8  `json:"-"` //    0x05
	GameDataFlag uint16 `json:"GameDataFlag"`//    0x06
	Score        uint32 `json:"Score"`//    0x08
	RealScore    uint32 `json:"RealScore"`//   Maybe
	Unk2         uint32 `json:"-"` //    0x10
}

type RequestLevelUpInfoPacket struct {
	PacketHead uint32 `json:"-"` //    0x00 0x0000001
	PacketType uint32 `json:"-"` //    0x04 0x100000C
	ProfileID  uint32 `json:"-"` //    0x08 0xBD5
}

type LevelUpInfoPacket struct {
	PacketHead uint32 `json:"-"` //    0x00 0x0000001
	PacketType uint32 `json:"-"` //    0x04 0x100000D
	ProfileID  uint32 //    0x08 0xBD5
	Level      uint32 //    0x0C
}

type GameOverPacket struct {
	PacketHead uint32 `json:"-"` //    0x00 0x0000001
	PacketType uint32 `json:"-"` //    0x04 0x1000001
	Unk0       uint32 `json:"-"` //    0x08 0xBD5
}

type WorldBestPacket struct {
	PacketHead  uint32 `json:"-"` //    0x00 0x00000002
	PacketType  uint32 `json:"-"` //    0x04 0x10000009
	Unk0        uint32 `json:"-"` //    0x08 5056
	Unk1        uint32 `json:"-"` //    0x0C 0x0000000F
	Unk2        uint32 `json:"-"` //    0x10 674200
	Unk3        uint32 `json:"-"` //    0x14 0x00000000
	Unk4        uint32 `json:"-"` //    0x18 0x00000000
	WorldScores [4095]WorldBestScore `json:"WorldScores"`
	Unk5        uint32 `json:"-"` //    0x?? 0x00000000
	Unk6        uint32 `json:"-"` //    0x?? 0x00000000
	PacketTrail uint32 `json:"-"` //    0x?? 0x00000000
}

type WorldBestScore struct {
	SongID     uint32      `json:"SongID"`//
	ChartLevel uint16      `json:"ChartLevel"`//
	ChartMode  uint16      `json:"ChartMode"`//
	Score      uint32      `json:"Score"`//
	Unk0       uint32      `json:"-"` //
	Unk1       uint32      `json:"-"` //
	Nickname   PIUNickname `json:"Nickname"`//
}

type RankModePacket struct {
	SimplePacket
	Ranks [400]SongRank `json:"Ranks"`
}

type SongRank struct {
	SongID uint32      `json:"SongID"`
	First  PIUNickname `json:"First"`
	Second PIUNickname `json:"Second"`
	Third  PIUNickname `json:"Third"`
}

var ACKPacketLength = int(binary.Size(ACKPacket{}))
var MachineInfoPacketLength = int(binary.Size(MachineInfoPacket{}))
var ScoreBoardPacketLength = int(binary.Size(ScoreBoardPacket{}))
var LoginPacketLength = int(binary.Size(LoginPacket{}))
var ProfilePacketLength = int(binary.Size(ProfilePacket{}))
var RequestLevelUpInfoPacketLength = int(binary.Size(RequestLevelUpInfoPacket{}))
var LevelUpInfoPacketLength = int(binary.Size(LevelUpInfoPacket{}))
var GameOverPacketLength = int(binary.Size(GameOverPacket{}))
var WorldBestPacketLength = int(binary.Size(WorldBestPacket{}))
var ProfileBusyPacketLength = int(binary.Size(ProfileBusyPacket{}))
var ByePacketLength = int(binary.Size(ByePacket{}))
var EnterProfilePacketLength = int(binary.Size(EnterProfilePacket{}))
var RequestWorldBestPacketLength = int(binary.Size(RequestWorldBestPacket{}))
var RankModePacketLength = int(binary.Size(RankModePacket{}))
var RequestRankModePacketLength = int(binary.Size(RequestRankModePacket{}))
var LoginPacketV2Length = int(binary.Size(LoginPacketV2{}))
var KeepAlivePacketLength = int(binary.Size(KeepAlivePacket{}))
var ScoreBoardPacket2Length = int(binary.Size(ScoreBoardPacket2{}))
var UnknownPacket0Length = int(binary.Size(UnknownPacket0{}))
var UnknownPacket1Length = int(binary.Size(UnknownPacket1{}))

var BiggestPacket = 0

func init() {

	fmt.Println(ScoreBoardPacket2Length)

	packetLens := []int{
		ACKPacketLength, MachineInfoPacketLength, MachineInfoPacketLength,
		ScoreBoardPacketLength, LoginPacketLength, ProfilePacketLength, RequestLevelUpInfoPacketLength,
		LevelUpInfoPacketLength, GameOverPacketLength, WorldBestPacketLength,
		ProfileBusyPacketLength, ByePacketLength, EnterProfilePacketLength,
		RequestWorldBestPacketLength, RankModePacketLength, RequestRankModePacketLength,
		LoginPacketV2Length, KeepAlivePacketLength, ScoreBoardPacket2Length,
		UnknownPacket0Length, UnknownPacket1Length,
	}

	for _, v := range packetLens {
		if v > BiggestPacket {
			BiggestPacket = v
		}
	}
}
