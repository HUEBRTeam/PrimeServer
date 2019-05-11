package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func (p *ACKPacket) FromBinary(data []byte) error {
	if len(data) != ACKPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", ACKPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *MachineInfoPacket) FromBinary(data []byte) error {
	if len(data) != MachineInfoPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", MachineInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ScoreBoardPacket) FromBinary(data []byte) error {
	if len(data) != ScoreBoardPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", ScoreBoardPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *LoginPacket) FromBinary(data []byte) error {
	if len(data) != LoginPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", LoginPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *ProfilePacket) FromBinary(data []byte) error {
	if len(data) != ProfilePacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", ProfilePacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *UScore) FromBinary(data []byte) error {
	if len(data) != UScoreLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", UScoreLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *RequestLevelUpInfoPacket) FromBinary(data []byte) error {
	if len(data) != RequestLevelUpInfoPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", RequestLevelUpInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *LevelUpInfoPacket) FromBinary(data []byte) error {
	if len(data) != LevelUpInfoPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", LevelUpInfoPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *GameOverPacket) FromBinary(data []byte) error {
	if len(data) != GameOverPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", GameOverPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *WorldBestPacket) FromBinary(data []byte) error {
	if len(data) != WorldBestPacketLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", WorldBestPacketLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}

func (p *WorldBestScore) FromBinary(data []byte) error {
	if len(data) != WorldBestScoreLength {
		return fmt.Errorf("expected payload to have %d bytes got %d instead", WorldBestScoreLength, len(data))
	}

	return binary.Read(bytes.NewReader(data), binary.LittleEndian, p)
}
