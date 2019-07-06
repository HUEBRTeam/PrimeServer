package proto

func MakeACKPacket() *ACKPacket {
	return &ACKPacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketACK,
		},
		PacketTrail: PacketTrail,
	}
}

func MakeProfileBusyPacket() *ProfileBusyPacket {
	return &ProfileBusyPacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketProfileBusy,
		},
		Unk0: 0,
	}
}

func MakeByePacket(profileId uint32) *ByePacket {
	return &ByePacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketBye,
		},
		ProfileID: profileId,
	}
}

func MakeEnterProfilePacket(profileId, machineId, playerId uint32) *EnterProfilePacket {
	return &EnterProfilePacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketEnterProfile,
		},
		ProfileID: profileId,
		PlayerID:  playerId,
		MachineID: machineId,
	}
}

func MakeRequestWorldBestPacket() *RequestWorldBestPacket {
	return &RequestWorldBestPacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketRequestWorldBest,
		},
	}
}

func MakeRequestRankMode() *RequestRankModePacket {
	return &RequestRankModePacket{
		SimplePacket: SimplePacket{
			PacketHead: 0x00000003,
			PacketType: PacketRequestRankMode,
		},
	}
}

func MakeWorldBestPacket(scores []WorldBestScore) *WorldBestPacket {
	s := [4095]WorldBestScore{}

	for i, v := range scores {
		if i < 4095 {
			s[i] = v
		}
	}

	return &WorldBestPacket{
		PacketHead:  0x00000002,
		PacketType:  PacketWorldBest,
		Unk0:        5056,
		Unk1:        0x0000000F,
		Unk2:        674200,
		Unk3:        0x00000000,
		Unk4:        0x00000000,
		WorldScores: s,
		Unk5:        0x00000000,
		Unk6:        0x00000000,
		PacketTrail: 0x00000000,
	}
}

func MakeRankModePacket(ranks []SongRank) *RankModePacket {
	r := [400]SongRank{}

	for i, v := range ranks {
		if i < 400 {
			r[i] = v
		}
	}

	return &RankModePacket{
		SimplePacket: SimplePacket{
			PacketHead: 0x00000003,
			PacketType: PacketRankMode,
		},
		Ranks: r,
	}
}

func MakeProfilePacketDefault(name string, accessCode string) *ProfilePacket {
	return &ProfilePacket{
		PacketHead:  0x0000001,
		PacketType:  PacketProfile,
		AccessCode:  MakePIUString32(accessCode),
		Nickname:    MakePIUNickName(name),
		CountryID:   uint8(196),
		Avatar:      uint8(41),
		Level:       0,
		EXP:         0,
		PP:          0,
		RankSingle:  0,
		RankDouble:  0,
		PlayCount:   0,
		Kcal:        0,
		Modifiers:   0,
		SpeedMod:    0,
		RunningStep: 0,
	}
}

func MakeProfilePacket(name string, country int, avatar int, modifiers int, speedmod int, accessCode string) *ProfilePacket {
	return &ProfilePacket{
		PacketHead:  0x0000001,
		PacketType:  PacketProfile,
		AccessCode:  MakePIUString32(accessCode),
		Nickname:    MakePIUNickName(name),
		CountryID:   uint8(country),
		Avatar:      uint8(avatar),
		Level:       0,
		EXP:         0,
		PP:          0,
		RankSingle:  0,
		RankDouble:  0,
		PlayCount:   0,
		Kcal:        0,
		Modifiers:   uint64(modifiers),
		SpeedMod:    uint32(speedmod),
		RunningStep: 0,
	}
}

func MakeLevelupInfoPacket(profileId, level uint32) *LevelUpInfoPacket {
	return &LevelUpInfoPacket{
		PacketHead: PacketHead,
		PacketType: PacketLevelUpInfo,
		ProfileID:  profileId,
		Level:      level,
	}
}

func MakeKeepAlivePacket() *KeepAlivePacket {
	return &KeepAlivePacket{
		SimplePacket: SimplePacket{
			PacketHead: PacketHead,
			PacketType: PacketKeepAlive,
		},
		PacketTrail: 65535,
	}
}
