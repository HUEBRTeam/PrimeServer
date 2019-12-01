package Storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/HUEBRTeam/PrimeServer/tools"
	"github.com/gofrs/uuid"
)

type DiskBackend struct {
	folder string
	mtx    sync.RWMutex
}

func MakeDiskBackend(folder string) *DiskBackend {
	_ = os.MkdirAll(folder, 0770)
	return &DiskBackend{
		folder: folder,
		mtx:    sync.RWMutex{},
	}
}

func (db *DiskBackend) GetProfile(accessCode string) (prof proto.ProfilePacket, err error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	if !db.profileExists(accessCode) {
		err = fmt.Errorf("no such profile")
		return
	}

	var data []byte

	data, err = ioutil.ReadFile(db.getProfilePath(accessCode))

	if err != nil {
		return
	}

	err = prof.FromBinary(data)

	return
}

func (db *DiskBackend) getProfilePath(accessCode string) string {
	return path.Join(db.folder, fmt.Sprintf("%s.primeprofile", accessCode))
}

func (db *DiskBackend) profileExists(accessCode string) bool {
	_, err := os.Stat(db.getProfilePath(accessCode))
	if err == nil {
		return true
	}

	return false
}

func (db *DiskBackend) ListProfiles() []os.FileInfo { // returns all files in profiles folder
	files, err := ioutil.ReadDir(db.folder)
	if err != nil {
		return []os.FileInfo{}
	}
	return files
}

func (db *DiskBackend) genAccessCode() string {
	u, _ := uuid.NewV4()
	id := strings.Replace(u.String(), "-", "", -1)

	if db.profileExists(id) {
		return db.genAccessCode() // Try again
	}

	return id
}

func (db *DiskBackend) CreateProfile(name string, country int, avatar int, modifiers int, noteskinspeed int) (profile proto.ProfilePacket, err error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	accessCode := db.genAccessCode()

	err = profile.FromBinary(proto.MakeProfilePacket(name, country, avatar, modifiers, noteskinspeed, accessCode).ToBinary())

	if err != nil {
		return
	}

	err = db.saveProfile(profile)

	return
}

func (db *DiskBackend) saveProfile(profile proto.ProfilePacket) error {

	accessCode := profile.AccessCode.String()

	return ioutil.WriteFile(db.getProfilePath(accessCode), profile.ToBinary(), 0770)
}

func (db *DiskBackend) saveWorldBest(wb *proto.WorldBestPacket) error {
	return ioutil.WriteFile("worldbest.bin", wb.ToBinary(), 0770)
}

func (db *DiskBackend) saveRankMode(rm *proto.RankModePacket) error {
	return ioutil.WriteFile("profile.bin", rm.ToBinary(), 0770)
}

func (db *DiskBackend) SaveProfile(profile proto.ProfilePacket) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	return db.saveProfile(profile)
}

func (db *DiskBackend) SaveWorldBest(wb *proto.WorldBestPacket) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	return db.saveWorldBest(wb)
}

func (db *DiskBackend) SaveRankMode(rm *proto.RankModePacket) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	return db.saveRankMode(rm)
}

func (db *DiskBackend) GetWorldBest() (wb *proto.WorldBestPacket, err error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	if !tools.IsFile("worldbest.bin") {
		err = fmt.Errorf("World Best Packet does not exist.")
		wb = proto.MakeWorldBestPacket(nil)
		return
	}

	var data []byte

	data, err = ioutil.ReadFile("worldbest.bin")

	if err != nil {
		return
	}

	err = wb.FromBinary(data)

	return
}

func (db *DiskBackend) GetRankMode() (rm *proto.RankModePacket, err error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	if !tools.IsFile("rankmode.bin") {
		err = fmt.Errorf("Rank Mode Packet does not exist.")
		rm = proto.MakeRankModePacket(nil)
		return
	}

	var data []byte

	data, err = ioutil.ReadFile("rankmode.bin")

	if err != nil {
		return
	}

	err = rm.FromBinary(data)

	return
}

func (db *DiskBackend) GetFolder() string {
	return db.folder
}
