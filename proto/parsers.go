package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

// region FromBinary
func (p *ACKPacket) FromBinary(data []byte) error {
	if len(data) != ACKPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ACKPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *MachineInfoPacket) FromBinary(data []byte) error {
	if len(data) != MachineInfoPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), MachineInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ScoreBoardPacket) FromBinary(data []byte) error {
	if len(data) != ScoreBoardPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ScoreBoardPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *LoginPacket) FromBinary(data []byte) error {
	if len(data) != LoginPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), LoginPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}
func (p *LoginPacketV2) FromBinary(data []byte) error {
	if len(data) != LoginPacketV2Length {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), LoginPacketV2Length, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ProfilePacket) FromBinary(data []byte) error {
	if len(data) != ProfilePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ProfilePacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *RequestLevelUpInfoPacket) FromBinary(data []byte) error {
	if len(data) != RequestLevelUpInfoPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), RequestLevelUpInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *LevelUpInfoPacket) FromBinary(data []byte) error {
	if len(data) != LevelUpInfoPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), LevelUpInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *GameOverPacket) FromBinary(data []byte) error {
	if len(data) != GameOverPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), GameOverPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *WorldBestPacket) FromBinary(data []byte) error {
	if len(data) != WorldBestPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), WorldBestPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ProfileBusyPacket) FromBinary(data []byte) error {
	if len(data) != ProfileBusyPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ProfileBusyPacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ByePacket) FromBinary(data []byte) error {
	if len(data) != ByePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ByePacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *EnterProfilePacket) FromBinary(data []byte) error {
	if len(data) != EnterProfilePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), EnterProfilePacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *RequestWorldBestPacket) FromBinary(data []byte) error {
	if len(data) != RequestWorldBestPacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), RequestWorldBestPacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *RankModePacket) FromBinary(data []byte) error {
	if len(data) != RankModePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), RankModePacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *RequestRankModePacket) FromBinary(data []byte) error {
	if len(data) != RequestRankModePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), RequestRankModePacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}
func (p *KeepAlivePacket) FromBinary(data []byte) error {
	if len(data) != KeepAlivePacketLength {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), KeepAlivePacketLength, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}
func (p *ScoreBoardPacket2) FromBinary(data []byte) error {
	if len(data) != ScoreBoardPacket2Length {
		return fmt.Errorf("(%s) expected payload to have %d bytes got %d instead", p.GetName(), ScoreBoardPacket2Length, len(data))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

// endregion
// region ToBinary
func (p *ACKPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *MachineInfoPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *ScoreBoardPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *LoginPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *LoginPacketV2) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *ProfilePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *UScore) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *RequestLevelUpInfoPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *LevelUpInfoPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *GameOverPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *WorldBestPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *WorldBestScore) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *ProfileBusyPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *ByePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *EnterProfilePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *RequestWorldBestPacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *RankModePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *RequestRankModePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *KeepAlivePacket) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

func (p *ScoreBoardPacket2) ToBinary() []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.LittleEndian, p)
	return b.Bytes()
}

// endregion
// region GetType
func (p *ACKPacket) GetType() uint32 {
	return PacketACK
}

func (p *MachineInfoPacket) GetType() uint32 {
	return PacketMachineInfo
}

func (p *ScoreBoardPacket) GetType() uint32 {
	return PacketScoreBoard
}

func (p *LoginPacket) GetType() uint32 {
	return PacketLogin
}

func (p *LoginPacketV2) GetType() uint32 {
	return PacketLoginV2
}

func (p *ProfilePacket) GetType() uint32 {
	return PacketProfile
}

func (p *RequestLevelUpInfoPacket) GetType() uint32 {
	return PacketRequestLevelUpInfo
}

func (p *LevelUpInfoPacket) GetType() uint32 {
	return PacketLevelUpInfo
}

func (p *GameOverPacket) GetType() uint32 {
	return PacketGameOver
}

func (p *WorldBestPacket) GetType() uint32 {
	return PacketWorldBest
}

func (p *ProfileBusyPacket) GetType() uint32 {
	return PacketProfileBusy
}

func (p *ByePacket) GetType() uint32 {
	return PacketBye
}

func (p *EnterProfilePacket) GetType() uint32 {
	return PacketEnterProfile
}

func (p *RequestWorldBestPacket) GetType() uint32 {
	return PacketRequestWorldBest
}

func (p *RankModePacket) GetType() uint32 {
	return PacketRankMode
}

func (p *RequestRankModePacket) GetType() uint32 {
	return PacketRequestRankMode
}

func (p *KeepAlivePacket) GetType() uint32 {
	return PacketKeepAlive
}

func (p *ScoreBoardPacket2) GetType() uint32 {
	return PacketKeepAlive
}

// endregion
// region GetName
func (p *ACKPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *MachineInfoPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *ScoreBoardPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *LoginPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *LoginPacketV2) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *ProfilePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *RequestLevelUpInfoPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *LevelUpInfoPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *GameOverPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *WorldBestPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *ProfileBusyPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *ByePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *EnterProfilePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *RequestWorldBestPacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *RankModePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *RequestRankModePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *KeepAlivePacket) GetName() string {
	return GetPacketName(p.GetType())
}

func (p *ScoreBoardPacket2) GetName() string {
	return GetPacketName(p.GetType())
}

// endregion

func DecodePacket(data []byte) (GenericPacket, error) {
	packetType := binary.LittleEndian.Uint32(data[4:8])
	var gp GenericPacket

	switch packetType {
	case PacketGameOver:
		gp = &GameOverPacket{}
	case PacketACK:
		gp = &ACKPacket{}
	case PacketLogin:
		gp = &LoginPacket{}
	case PacketProfile:
		gp = &ProfilePacket{}
	case PacketProfileBusy:
		gp = &ProfileBusyPacket{}
	case PacketRequestWorldBest:
		gp = &RequestWorldBestPacket{}
	case PacketWorldBest:
		gp = &WorldBestPacket{}
	case PacketRequestRankMode:
		gp = &RequestRankModePacket{}
	case PacketRankMode:
		gp = &RankModePacket{}
	case PacketRequestLevelUpInfo:
		gp = &RequestLevelUpInfoPacket{}
	case PacketLevelUpInfo:
		gp = &LevelUpInfoPacket{}
	case PacketScoreBoard:
		gp = &ScoreBoardPacket{}
	case PacketEnterProfile:
		gp = &EnterProfilePacket{}
	case PacketBye:
		gp = &ByePacket{}
	case PacketMachineInfo:
		gp = &MachineInfoPacket{}
	case PacketLoginV2:
		gp = &LoginPacketV2{}
	case PacketKeepAlive:
		gp = &KeepAlivePacket{}
	case PacketScoreBoardV2:
		gp = &ScoreBoardPacket2{}
	default:
		_ = ioutil.WriteFile(fmt.Sprintf("%08x.bin", packetType), data, 0777)
		return nil, fmt.Errorf("no such packet type %08x", packetType)
	}

	err := gp.FromBinary(data)

	if err != nil {
		return nil, err
	}

	return gp, nil
}
