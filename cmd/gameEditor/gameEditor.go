package main

import (
	"github.com/quan-to/slog"
	"io/ioutil"
)

var log = slog.Scope("GameEditor")

func main() {
	data, err := ioutil.ReadFile("piu_108o_original")
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

	if !SetPrimeServerIP(data, "8.8.8.8") {
		log.Fatal("Your IP does not fit into Prime memory :(")
	}

	o, ip = FindPrimeServerIP(data)

	log.Info("IP Found at %d: %s", o, ip)
}
