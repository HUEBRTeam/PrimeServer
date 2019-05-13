package ProfileManager

import (
	"fmt"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/quan-to/slog"
	"sync"
	"time"
)

var log = slog.Scope("ProfileManager")

type ProfileManager struct {
	sb                    ProfileStorageBackend
	loadedProfiles        map[string]ProfileSession
	profileIdCount        uint32
	profileIdToAccessCode map[uint32]string
	mtx                   sync.Mutex
}

func MakeProfileManager(sb ProfileStorageBackend) *ProfileManager {
	return &ProfileManager{
		sb:                    sb,
		loadedProfiles:        map[string]ProfileSession{},
		profileIdToAccessCode: map[uint32]string{},
		mtx:                   sync.Mutex{},
		profileIdCount:        0,
	}
}

func (pm *ProfileManager) Create(name string) (string, error) {
	p, err := pm.sb.CreateProfile(name)

	if err != nil {
		return "", err
	}

	return p.AccessCode.String(), nil
}

func (pm *ProfileManager) Load(accessCode string, playerId uint32) (profile proto.ProfilePacket, err error) {
	profile, err = pm.sb.GetProfile(accessCode)
	if err != nil {
		return
	}

	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	_, ok := pm.loadedProfiles[accessCode]
	if ok {
		err = fmt.Errorf("profile busy")
		return
	}

	profile.ProfileID = pm.profileIdCount
	profile.PlayerID = playerId

	pm.profileIdToAccessCode[profile.ProfileID] = accessCode
	pm.profileIdCount++

	log.Debug("Loaded profile %s with ID %d", profile.Nickname, profile.ProfileID)

	return
}

func (pm *ProfileManager) GetProfileNickname(profileId uint32) string {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[profileId]

	if !ok {
		return "Unknown"
	}

	if v, ok := pm.loadedProfiles[accessCode]; ok {
		return v.Profile.Nickname.String()
	}

	return "Unknown"
}

func (pm *ProfileManager) LockProfile(profileId uint32) error {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[profileId]

	if !ok {
		return fmt.Errorf("profile busy")
	}

	profile, err := pm.sb.GetProfile(accessCode)
	if err != nil {
		return err
	}

	_, ok = pm.loadedProfiles[accessCode]
	if ok {
		return fmt.Errorf("profile busy")
	}

	log.Debug("User %s now played %d times. Saving profile...", profile.Nickname, profile.PlayCount)
	profile.PlayCount++
	err = pm.sb.SaveProfile(profile)
	if err != nil {
		log.Error("Error saving %s profile: %s", profile.Nickname, err)
	}

	pm.loadedProfiles[accessCode] = ProfileSession{
		Profile:    profile,
		LastUpdate: time.Now(),
	}

	return nil
}

func (pm *ProfileManager) KeepAlive(profileId uint32) {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[profileId]

	if !ok {
		return
	}

	if v, ok := pm.loadedProfiles[accessCode]; ok {
		v.LastUpdate = time.Now()
		pm.loadedProfiles[accessCode] = v
	}
}

func (pm *ProfileManager) Unload(profileId uint32) {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[profileId]

	if !ok {
		return
	}

	v, ok := pm.loadedProfiles[accessCode]
	if ok {
		delete(pm.loadedProfiles, accessCode)
	}

	err := pm.sb.SaveProfile(v.Profile)
	if err != nil {
		log.Error("Error saving %s profile: %s", v.Profile.Nickname, err)
	}
}

func (pm *ProfileManager) PutScoreBoard(sb proto.ScoreBoardPacket) {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[sb.ProfileID]

	if !ok {
		return
	}

	v, ok := pm.loadedProfiles[accessCode]
	if ok {
		p := v.Profile

		p.EXP += uint64(sb.EXP)
		p.PP += uint64(sb.PP)
		p.RunningStep += uint64(sb.RunningStep)
		p.Kcal += sb.Kcal

		v.Profile = p
		err := pm.sb.SaveProfile(p)
		if err != nil {
			log.Error("Error saving %s profile: %s", p.Nickname, err)
		}
	}
}

func (pm *ProfileManager) PutScoreBoard2(sb proto.ScoreBoardPacket2) {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	accessCode, ok := pm.profileIdToAccessCode[sb.ProfileID]

	if !ok {
		return
	}

	v, ok := pm.loadedProfiles[accessCode]
	if ok {
		p := v.Profile

		p.EXP += uint64(sb.EXP)
		p.PP += uint64(sb.PP)
		p.RunningStep += uint64(sb.RunningStep)
		p.Kcal += sb.Kcal

		v.Profile = p
		err := pm.sb.SaveProfile(p)
		if err != nil {
			log.Error("Error saving %s profile: %s", p.Nickname, err)
		}
	}
}
