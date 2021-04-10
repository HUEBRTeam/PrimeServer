package ProfileManager

import (
	"github.com/HUEBRTeam/PrimeServer/proto"
	"time"
)

type ProfileSession struct {
	Profile    proto.ProfilePacket
	LastUpdate time.Time
}
