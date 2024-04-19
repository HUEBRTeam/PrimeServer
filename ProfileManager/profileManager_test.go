package ProfileManager

import (
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/stretchr/testify/require"
	"testing"
)

type FakeProfileStorageBackend struct {
}

func (ps *FakeProfileStorageBackend) GetProfile(accessCode string) (proto.ProfilePacket, error) {
	return proto.ProfilePacket{}, nil
}
func (ps *FakeProfileStorageBackend) CreateProfile(name string, country int, avatar int, modifiers int, noteskinspeed int) (proto.ProfilePacket, error) {
	return proto.ProfilePacket{}, nil
}
func (ps *FakeProfileStorageBackend) SaveProfile(packet proto.ProfilePacket) error  { return nil }
func (ps *FakeProfileStorageBackend) SaveWorldBest(wb *proto.WorldBestPacket) error { return nil }
func (ps *FakeProfileStorageBackend) SaveRankMode(rm *proto.RankModePacket) error   { return nil }
func (ps *FakeProfileStorageBackend) GetWorldBest() (wb *proto.WorldBestPacket, err error) {
	return &proto.WorldBestPacket{}, nil
}
func (ps *FakeProfileStorageBackend) GetRankMode() (rm *proto.RankModePacket, err error) {
	return &proto.RankModePacket{}, nil
}
func (ps *FakeProfileStorageBackend) GetFolder() string { return "" }

func TestProfileManager_PutScoreBoard2(t *testing.T) {
	const accessCode = "fake-access-code"
	profile := proto.ProfilePacket{
		ProfileID: uint32(1),
	}
	profile.Scores[0] = proto.UScore{
		SongID:     123,
		ChartLevel: 1,
		Score:      5000,
		Unk2:       1,
	}

	pm := &ProfileManager{
		sb: &FakeProfileStorageBackend{},
		profileIdToAccessCode: map[uint32]string{
			profile.ProfileID: accessCode,
		},
		loadedProfiles: map[string]ProfileSession{
			"fake-access-code": {
				Profile: profile,
			},
		},
	}

	t.Run("when profile plays a new song, add a new score entry", func(t *testing.T) {
		pm.loadedProfiles[accessCode] = ProfileSession{
			Profile: profile,
		}
		pm.PutScoreBoard2(proto.ScoreBoardPacket2{
			ProfileID:  profile.ProfileID,
			SongID:     456,
			ChartLevel: 2,
			Score:      10000,
		})

		p := pm.loadedProfiles[accessCode].Profile
		t.Run("existing scores remain unchanged", func(t *testing.T) {
			require.Equal(t, uint32(123), p.Scores[0].SongID)
			require.Equal(t, uint8(1), p.Scores[0].ChartLevel)
			require.Equal(t, uint32(5000), p.Scores[0].Score)
		})
		t.Run("add a new score entry", func(t *testing.T) {
			require.Equal(t, uint32(456), p.Scores[1].SongID)
			require.Equal(t, uint8(2), p.Scores[1].ChartLevel)
			require.Equal(t, uint32(10000), p.Scores[1].Score)
			require.Equal(t, uint32(1), p.Scores[1].Unk2)
		})
		t.Run("do not overflow", func(t *testing.T) {
			require.Equal(t, uint32(0), p.Scores[2].SongID)
			require.Equal(t, uint8(0), p.Scores[2].ChartLevel)
			require.Equal(t, uint32(0), p.Scores[2].Score)
		})
	})
	t.Run("when profile plays the same song but on a different chart level", func(t *testing.T) {
		pm.loadedProfiles[accessCode] = ProfileSession{
			Profile: profile,
		}
		pm.PutScoreBoard2(proto.ScoreBoardPacket2{
			ProfileID:  profile.ProfileID,
			SongID:     123,
			ChartLevel: 2,
			Score:      10000,
		})

		p := pm.loadedProfiles[accessCode].Profile
		t.Run("existing scores remain unchanged", func(t *testing.T) {
			require.Equal(t, uint32(123), p.Scores[0].SongID)
			require.Equal(t, uint8(1), p.Scores[0].ChartLevel)
			require.Equal(t, uint32(5000), p.Scores[0].Score)
		})
		t.Run("add a new score entry", func(t *testing.T) {
			require.Equal(t, uint32(123), p.Scores[1].SongID)
			require.Equal(t, uint8(2), p.Scores[1].ChartLevel)
			require.Equal(t, uint32(10000), p.Scores[1].Score)
			require.Equal(t, uint32(1), p.Scores[1].Unk2)
		})
		t.Run("do not overflow", func(t *testing.T) {
			require.Equal(t, uint32(0), p.Scores[2].SongID)
			require.Equal(t, uint8(0), p.Scores[2].ChartLevel)
			require.Equal(t, uint32(0), p.Scores[2].Score)
		})
	})
	t.Run("when profile plays the same song and same chart level with a higher score", func(t *testing.T) {
		pm.loadedProfiles[accessCode] = ProfileSession{
			Profile: profile,
		}
		pm.PutScoreBoard2(proto.ScoreBoardPacket2{
			ProfileID:  profile.ProfileID,
			SongID:     123,
			ChartLevel: 1,
			Score:      10000,
		})

		p := pm.loadedProfiles[accessCode].Profile
		t.Run("replace the existing score with the new high score", func(t *testing.T) {
			require.Equal(t, uint32(123), p.Scores[0].SongID)
			require.Equal(t, uint8(1), p.Scores[0].ChartLevel)
			require.Equal(t, uint32(10000), p.Scores[0].Score)
		})
		t.Run("do not overflow", func(t *testing.T) {
			require.Equal(t, uint32(0), p.Scores[1].SongID)
			require.Equal(t, uint8(0), p.Scores[1].ChartLevel)
			require.Equal(t, uint32(0), p.Scores[1].Score)
		})
	})
	t.Run("when profile plays the same song and same chart level with a lower score", func(t *testing.T) {
		pm.loadedProfiles[accessCode] = ProfileSession{
			Profile: profile,
		}
		pm.PutScoreBoard2(proto.ScoreBoardPacket2{
			ProfileID:  profile.ProfileID,
			SongID:     123,
			ChartLevel: 1,
			Score:      2000,
		})

		p := pm.loadedProfiles[accessCode].Profile
		t.Run("do not register it to maintain the existing higher score", func(t *testing.T) {
			require.Equal(t, uint32(123), p.Scores[0].SongID)
			require.Equal(t, uint8(1), p.Scores[0].ChartLevel)
			require.Equal(t, uint32(5000), p.Scores[0].Score)
		})
		t.Run("do not overflow", func(t *testing.T) {
			require.Equal(t, uint32(0), p.Scores[1].SongID)
			require.Equal(t, uint8(0), p.Scores[1].ChartLevel)
			require.Equal(t, uint32(0), p.Scores[1].Score)
		})
	})
}
