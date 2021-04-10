package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"github.com/HUEBRTeam/PrimeServer"
	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/Storage"
	"github.com/HUEBRTeam/PrimeServer/cmd/server/network"
	"github.com/HUEBRTeam/PrimeServer/cmd/server/rest"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/HUEBRTeam/PrimeServer/tools"
	"github.com/quan-to/slog"
)

const (
	ConnPort   = "60010"
	ConnType   = "tcp"
	ConfigFile = "config.json"
	ProfileDir = "profiles/"
)

var log = slog.Scope("PrimeServer")
var profileManager *ProfileManager.ProfileManager
var config = Config{}

func main() {
	sb := Storage.MakeDiskBackend("profiles")
	profileManager = ProfileManager.MakeProfileManager(sb)
	if !tools.IsFile(ConfigFile) {
		log.Info("Config file not found, creating one...")
		j, err := json.Marshal(Config{})
		if err != nil {
			log.Error("Error: cannot read config struct, %s", err.Error()) // this shouldn't happen
			os.Exit(1)
		}
		err = ioutil.WriteFile(ConfigFile, []byte(j), 0644)
		if err != nil {
			log.Error("Error: cannot write config file, %s", err.Error())
			os.Exit(1)
		}
	}
	conf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Error("Error: cannot read config file, %s", err.Error())
		os.Exit(1)
	}
	_ = json.Unmarshal(conf, &config)
	if config.Online {
		if tools.IsDir(sb.GetFolder()) {
			for _, f := range sb.ListProfiles() {
				prof, err := network.RetrieveProfile(config.APIKey, strings.Replace(f.Name(), ".primeprofile", "", -1), config.ServerAddress, profileManager)
				if err != nil {
					log.Error("Error: could not retrieve profile for access code %s, skipping... %s", strings.Replace(f.Name(), ".primeprofile", "", -1), err.Error())
					break
				}
				sb.SaveProfile(prof)
			}
		}
	}
	rs := rest.MakeRestServer(8090, profileManager)

	go func() {
		log.Fatal(rs.Listen())
	}()

	// Listen for incoming connections.
	l, err := net.Listen(ConnType, ":"+ConnPort)
	if err != nil {
		log.Error("Error listening: %s", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	log.Info("Listening on :%s", ConnPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Error("Error accepting: %s", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handlePacket(cs *PrimeServer.ClientSession, packet []byte) {
	cs.Lock()
	defer cs.Unlock()

	l := slog.Scope(cs.Conn.RemoteAddr().String())
	dec, ok := proto.DecryptPacket(packet)

	if !ok {
		l.Error("Received invalid packet.")
		return
	}

	p, err := proto.DecodePacket(dec)

	if err != nil {
		l.Error("Error parsing packet: %s", err)
		return
	}

	switch v := p.(type) {
	case *proto.RequestWorldBestPacket:
		handleRequestWorldBestPacket(l, cs.Conn, *v)

	case *proto.RequestRankModePacket:
		handleRequestRankModePacket(l, cs.Conn, *v)

	case *proto.LoginPacket:
		handleLoginPacket(l, cs.Conn, *v)

	case *proto.LoginPacketV2:
		handleLoginPacketV2(l, cs.Conn, *v)

	case *proto.EnterProfilePacket:
		handleEnterProfilePacket(l, cs.Conn, *v)

	case *proto.MachineInfoPacket:
		handleMachineInfoPacket(l, cs.Conn, *v)

	case *proto.ScoreBoardPacket:
		handleScoreBoardPacket(l, cs.Conn, *v)

	case *proto.ScoreBoardPacket2:
		handleScoreBoardV2Packet(l, cs.Conn, *v)

	case *proto.RequestLevelUpInfoPacket:
		handleRequestLevelupInfo(l, cs.Conn, *v)

	case *proto.ByePacket:
		handleByePacket(l, cs.Conn, *v)

	case *proto.KeepAlivePacket:
		PrimeServer.SendACK(cs.Conn)

	default:
		l.Debug("Received Packet: %s", p.GetName())
		PrimeServer.SendACK(cs.Conn)
	}
}

func handleKeepAlive(cs *PrimeServer.ClientSession) {
	for cs.Running {
		time.Sleep(time.Second * 5)
		cs.Lock()
		PrimeServer.SendKeepAlive(cs.Conn)
		cs.Unlock()
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	l := slog.Scope(conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	c := 0
	defer conn.Close()

	log.Debug("Connection from %s", conn.RemoteAddr().String())

	cs := PrimeServer.MakeClientSession(conn)
	cs.Running = true

	go handleKeepAlive(cs)

	for cs.Running {
		// Read the incoming connection into the buffer.
		n, err := conn.Read(buf[c:])
		if err != nil {
			if err != io.EOF {
				l.Error("Error reading: %s", err.Error())
			} else {
				l.Info("Client disconnected")
			}
			cs.Running = false
			return
		}
		c += n

		if c > 4 {
			plen := int(binary.LittleEndian.Uint32(buf[:4]))

			if plen > proto.BiggestPacket {
				l.Error("Invalid packet length: %d", plen)
				cs.Running = false
				return
			}

			if c >= plen {
				handlePacket(cs, buf[4:plen])
				copy(buf, buf[plen:])
				c -= plen
			}
		}
	}
}
