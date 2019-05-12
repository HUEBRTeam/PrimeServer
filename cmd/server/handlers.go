package main

import (
	"github.com/HUEBRTeam/PrimeServer"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/quan-to/slog"
	"net"
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
	l.Debug(v)
	PrimeServer.SendACK(conn)
}

func handleRequestWorldBestPacket(l *slog.Instance, conn net.Conn, v proto.RequestWorldBestPacket) {
	wb := proto.MakeWorldBestPacket(nil)
	PrimeServer.SendPacket(conn, wb.ToBinary())
}

func handleRequestRankModePacket(l *slog.Instance, conn net.Conn, v proto.RequestRankModePacket) {
	rm := proto.MakeRankModePacket(nil)
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

func handleByePacket(l *slog.Instance, conn net.Conn, v proto.ByePacket) {
	nick := profileManager.GetProfileNickname(v.ProfileID)

	l.Info("User %s finished playing.", nick)

	profileManager.Unload(v.ProfileID)
	PrimeServer.SendACK(conn)
}
