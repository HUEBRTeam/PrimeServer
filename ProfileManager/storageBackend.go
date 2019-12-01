package ProfileManager

import "github.com/HUEBRTeam/PrimeServer/proto"

type ProfileStorageBackend interface {
	GetProfile(accessCode string) (proto.ProfilePacket, error)
	CreateProfile(name string, country int, avatar int, modifiers int, noteskinspeed int) (proto.ProfilePacket, error)
	SaveProfile(packet proto.ProfilePacket) error
	SaveWorldBest(wb *proto.WorldBestPacket) error
	SaveRankMode(rm *proto.RankModePacket) error
	GetWorldBest() (wb *proto.WorldBestPacket, err error)
	GetRankMode() (rm *proto.RankModePacket, err error)
	GetFolder() string
}
