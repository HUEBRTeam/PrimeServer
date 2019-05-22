package main

import (
	"encoding/hex"
	"fmt"
	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/Storage"
	"github.com/quan-to/slog"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
)

var log = slog.Scope("Profile Maker")

func main() {
	name := kingpin.Arg("name", "The name of the profile").Required().String()
	kingpin.Parse()

	if len(*name) > 12 {
		log.Fatal("The maximum length of name is 12. The name \"%s\" has %d characters.", *name, len(*name))
	}

	log.Info("Initializing Profile Manager")
	sb := Storage.MakeDiskBackend("profiles")
	pm := ProfileManager.MakeProfileManager(sb)

	log.Info("Creating profile %s", *name)
	accessCode, err := pm.Create(*name)

	if err != nil {
		log.Fatal(err)
	}

	v, _ := hex.DecodeString(accessCode)

	profileFile := fmt.Sprintf("prime-%s.bin", *name)

	log.Info("Saving file %s", profileFile)
	err = ioutil.WriteFile(profileFile, v, 0777)

	if err != nil {
		log.Fatal("Error saving file %s: %s", profileFile, err)
	}

	log.Info("Your access code to profile %s is %s. I created a %s file to use in USB device.", *name, accessCode, profileFile)
	log.Info("Put the profiles/%s.primeprofile file into your profiles server folder.", accessCode)
	log.Info("Have fun!")
}
