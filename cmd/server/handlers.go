package main

import (
	"net"
	"fmt"

	"github.com/HUEBRTeam/PrimeServer"
	"github.com/HUEBRTeam/PrimeServer/cmd/server/network"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/quan-to/slog"
)

func handleLoginPacket(l *slog.Instance, conn net.Conn, v proto.LoginPacket) {
	ac := v.AccessCode.String()
	p, err := profileManager.Load(ac, v.PlayerID)
	if err != nil {
		l.Error("Error loading profile %s: %s", ac, err)
		PrimeServer.SendProfileBusy(conn)
	} else {
		PrimeServer.SendProfile(conn, p)
	}
}

func handleLoginPacketV2(l *slog.Instance, conn net.Conn, v proto.LoginPacketV2) {
	ac := v.AccessCode.String()
	if config.Online {
		prof, err := network.RetrieveProfile(config.APIKey, ac, config.ServerAddress, profileManager)
		if err != nil {
			log.Error("Error: could not retrieve profile for access code %s, skipping... %s", ac, err.Error())
		} else {
			err = profileManager.GetStorageBackend().SaveProfile(prof)
			if err != nil {
				log.Error("Error: could not save profile %s", err.Error())
			}
		}
		/*if md5check { // probably only do this for score submissions
			calculate packets md5
			try connecting with apikey to retrieve md5 of world best and rank mode
			if error break, log and just skip this stuff
		}*/
		wb, err := network.RetrieveWorldBest(config.APIKey, config.ServerAddress, config.ScoreType)
		if err != nil {
			log.Error("Error: could not retrieve World Best packet %s", err.Error())
		} else {
			err = profileManager.GetStorageBackend().SaveWorldBest(wb)
			if err != nil {
				log.Error("Error: could not save World Best packet %s", err.Error())
			}
		}
		rank, err := network.RetrieveRankMode(config.APIKey, config.ServerAddress, config.ScoreType)
		if err != nil {
			log.Error("Error: could not retrieve Rank Mode packet %s", err.Error())
		} else {
			fmt.Printf("%+v\n", rank)
			err = profileManager.GetStorageBackend().SaveRankMode(rank)
			if err != nil {
				log.Error("Error: could not save Rank Mode packet %s", err.Error())
			}
		}

	}
	p, err := profileManager.Load(ac, v.PlayerID)
	if err != nil {
		l.Error("Error loading profile %s: %s", ac, err)
		PrimeServer.SendProfileBusy(conn)
	} else {
		PrimeServer.SendProfile(conn, p)
	}
}

func handleEnterProfilePacket(l *slog.Instance, conn net.Conn, v proto.EnterProfilePacket) {
	err := profileManager.LockProfile(v.ProfileID)
	if err != nil {
		l.Error("Error locking profile %s: %s", v.ProfileID, err)
		PrimeServer.SendProfileBusy(conn)
	} else {
		PrimeServer.SendACK(conn)
	}
}

func handleMachineInfoPacket(l *slog.Instance, conn net.Conn, v proto.MachineInfoPacket) {
	l.Debug(v.String())
	if config.Online {
		err := network.SubmitMachineInfo(config.APIKey, config.ServerAddress, v)
		if err != nil {
			log.Error("Error: could not send Machine ID packet %s", err.Error())
		}
	}
	PrimeServer.SendACK(conn)
}

func handleRequestWorldBestPacket(l *slog.Instance, conn net.Conn, v proto.RequestWorldBestPacket) {
	wb, err := profileManager.GetStorageBackend().GetWorldBest()
	if err != nil {
		log.Error("Error: could not get World Best packet %s", err.Error())
	}
	PrimeServer.SendPacket(conn, wb.ToBinary())
}

func handleRequestRankModePacket(l *slog.Instance, conn net.Conn, v proto.RequestRankModePacket) {
	rm, err := profileManager.GetStorageBackend().GetRankMode()
	if err != nil {
		log.Error("Error: could not get Rank Mode packet %s", err.Error())
	}
	PrimeServer.SendPacket(conn, rm.ToBinary())
}

func handleRequestLevelupInfo(l *slog.Instance, conn net.Conn, v proto.RequestLevelUpInfoPacket) {
	profileManager.KeepAlive(v.ProfileID)
	PrimeServer.SendLevelUpPacket(conn, v.ProfileID, 0)
}

func handleScoreBoardPacket(l *slog.Instance, conn net.Conn, v proto.ScoreBoardPacket) {
	nick := profileManager.GetProfileNickname(v.ProfileID)
	profileManager.KeepAlive(v.ProfileID)
	profileManager.PutScoreBoard(v)
	l.Info("User %s played a song", nick)
	PrimeServer.SendACK(conn)
}

func handleScoreBoardV2Packet(l *slog.Instance, conn net.Conn, v proto.ScoreBoardPacket2) {
	nick := profileManager.GetProfileNickname(v.ProfileID)
	profileManager.KeepAlive(v.ProfileID)
	profileManager.PutScoreBoard2(v)
	l.Info("User %s played a song", nick)
	l.Info("%s", v.String())
	if config.Online {
		err := network.SubmitScore(config.APIKey, config.ServerAddress, v, profileManager.ProfileIDToAccessCode(v.ProfileID))
		if err != nil {
			log.Error("Error: could not send Score packet %s", err.Error())
		}
		p, err := profileManager.GetStorageBackend().GetProfile(profileManager.ProfileIDToAccessCode(v.ProfileID))
		if err != nil {
			log.Error("Error: could not get profile %s", err.Error())
		}
		err = network.SubmitProfile(config.APIKey, config.ServerAddress, p, profileManager.ProfileIDToAccessCode(v.ProfileID))
		if err != nil {
			log.Error("Error: could not send Profile packet %s", err.Error())
		}
		/*if md5check { // probably only do this for score submissions
			calculate packets md5
			try connecting with apikey to retrieve md5 of world best and rank mode
			if error break, log and just skip this stuff
		}*/
		wb, err := network.RetrieveWorldBest(config.APIKey, config.ServerAddress, config.ScoreType)
		if err != nil {
			log.Error("Error: could not retrieve World Best packet %s", err.Error())
		} else {
			err = profileManager.GetStorageBackend().SaveWorldBest(wb)
			if err != nil {
				log.Error("Error: could not save World Best packet %s", err.Error())
			}
		}
		rank, err := network.RetrieveRankMode(config.APIKey, config.ServerAddress, config.ScoreType)
		if err != nil {
			log.Error("Error: could not retrieve Rank Mode packet %s", err.Error())
		} else {
			err = profileManager.GetStorageBackend().SaveRankMode(rank)
			if err != nil {
				log.Error("Error: could not save Rank Mode packet %s", err.Error())
			}
		}
	}
	PrimeServer.SendACK(conn)
}

func handleByePacket(l *slog.Instance, conn net.Conn, v proto.ByePacket) {
	nick := profileManager.GetProfileNickname(v.ProfileID)

	l.Info("User %s finished playing.", nick)

	profileManager.Unload(v.ProfileID)
	PrimeServer.SendACK(conn)
}
