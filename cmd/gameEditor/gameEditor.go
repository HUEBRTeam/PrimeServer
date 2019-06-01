package main

import (
	"bufio"
	"github.com/HUEBRTeam/PrimeServer/tools"
	"github.com/quan-to/slog"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var log = slog.Scope("GameEditor")

func main() {
	reader := bufio.NewReader(os.Stdin)

	log.Info("Please type the filename to edit: ")

	filename, err := reader.ReadString('\n')

	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
		os.Exit(1)
	}

	filename = strings.Trim(filename, "\r\n ")

	if !tools.Exists(filename) {
		log.Fatal("File %s does not exists", filename)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	version, err := GetGameVersion(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Game Version: %s", version)

	o := FindPrimeDaemonURL(data)

	log.Info("Prime Daemon Offset: %d", o)

	o, ip := FindPrimeServerIP(data)

	log.Info("IP Found at %d: %s", o, ip)

	sl := true
	for sl {
		log.Info("Please enter the IP you want to write: ")
		ip, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			os.Exit(1)
		}

		ip = strings.Trim(ip, "\r\n ")

		if !tools.IsValidIP4(ip) {
			log.Error("The IP Address \"%s\" is not valid.", ip)
			continue
		}

		log.Info("Do you want to set the IP %s to game?", ip)

		r, _ := reader.ReadString('\n')

		r = strings.ToLower(r)

		if len(r) == 0 || (r[0] != 's' && r[0] != 'y') {
			continue
		}

		if !SetPrimeServerIP(data, ip) {
			log.Error("Your IP does not fit into Prime memory :(")
			continue
		}

		o, ip = FindPrimeServerIP(data)

		log.Info("Current IP: %s", ip)
		log.Info("I will save the file. That's correct?")

		r, _ = reader.ReadString('\n')
		r = strings.ToLower(r)

		if len(r) == 0 || (r[0] != 's' && r[0] != 'y') {
			continue
		}

		err = ioutil.WriteFile(filename, data, os.ModePerm)
		if err != nil {
			log.Error(err)
			continue
		}

		sl = false
	}

	log.Info("Done!")
}
