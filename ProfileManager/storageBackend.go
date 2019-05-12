package ProfileManager

import "github.com/HUEBRTeam/PrimeServer/proto"

type ProfileStorageBackend interface {
	GetProfile(accessCode string) (proto.ProfilePacket, error)
	CreateProfile(name string) (proto.ProfilePacket, error)
	SaveProfile(packet proto.ProfilePacket) error
}
