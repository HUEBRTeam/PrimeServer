package ProfileManager

import (
	"fmt"
	"sync"
	"time"

	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/quan-to/slog"
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

func (pm *ProfileManager) Create(name string, country int, avatar int, modifiers int, noteskinspeed int) (string, error) {
	p, err := pm.sb.CreateProfile(name, country, avatar, modifiers, noteskinspeed)

	if err != nil {
		return "", err
	}

	return p.AccessCode.String(), nil
}

func (pm *ProfileManager) Change(accessCode string, name string, country int, avatar int, modifiers int, noteskinspeed int) error {
	pm.mtx.Lock()
	defer pm.mtx.Unlock()

	profile, err := pm.sb.GetProfile(accessCode)
	if err != nil {
		log.Error("Error getting %s profile: %s", profile.Nickname, err)
		return err
	}

	profile.Nickname = proto.MakePIUNickName(name)
	profile.CountryID = uint8(country)
	profile.Avatar = uint8(avatar)
	profile.Modifiers = uint64(modifiers)
	profile.NoteSkinSpeed = uint32(noteskinspeed)

	err = pm.sb.SaveProfile(profile)
	if err != nil {
		log.Error("Error saving %s profile: %s", profile.Nickname, err)
		return err
	}

	return nil
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

	log.Debug("Saved profile %s", v.Profile.Nickname)
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
		pm.loadedProfiles[accessCode] = v
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
		p.Modifiers = uint64(sb.Modifiers)
		p.NoteSkinSpeed = sb.NoteSkinSpeed
		p.RushSpeed = sb.RushSpeed
		for i := 0; i < len(p.Scores); i++ {
			score := &p.Scores[i]
			if score.SongID != 0 {
				if score.SongID != sb.SongID {
					continue
				}
				if score.ChartLevel != uint8(sb.ChartLevel) {
					continue
				}
			}
			if score.Score <= sb.Score {
				score.SongID = sb.SongID
				score.ChartLevel = uint8(sb.ChartLevel)
				score.Unk0 = 0
				score.GameDataFlag = 0
				score.Score = sb.Score
				score.RealScore = sb.Score
				score.Unk2 = 1
			}
			break
		}

		v.Profile = p
		pm.loadedProfiles[accessCode] = v
		err := pm.sb.SaveProfile(p)
		if err != nil {
			log.Error("Error saving %s profile: %s", p.Nickname, err)
		}
	}
}

func (pm *ProfileManager) GetStorageBackend() ProfileStorageBackend {
	return pm.sb
}

func (pm *ProfileManager) ProfileIDToAccessCode(id uint32) string {
	return pm.profileIdToAccessCode[id]
}
