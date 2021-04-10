package proto

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestMachineInfoPacket_FromBinary(t *testing.T) {
	data, err := ioutil.ReadFile("../RE/Packets/Start/1.01.0/raw_8_out_115.68.108.183_60000.bin")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	dec, ok := DecryptPacket(data[4:])

	if !ok {
		t.Fatalf("Error decrypting packet")
	}

	mip := MachineInfoPacket{}
	err = mip.FromBinary(dec)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}

	t.Log(mip.String())
}

func TestProfilePacket_FromBinary(t *testing.T) {
	data, err := ioutil.ReadFile("../RE/Packets/Login/sodium_12_in_pkt.bin")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	pp := ProfilePacket{}
	err = pp.FromBinary(data)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}
}

func TestScoreBoard2_FromBinary(t *testing.T) {
	data, err := ioutil.ReadFile("../RE/Packets/MusicScore/1.08.0/PUMP IT UP: PRIME_224.bin")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	pp := ScoreBoardPacket2{}
	err = pp.FromBinary(data)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}

	fmt.Println(pp.String())

	data, err = ioutil.ReadFile("../RE/Packets/MusicScore/1.08.0/PUMP IT UP: PRIME_225.bin")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	pp = ScoreBoardPacket2{}
	err = pp.FromBinary(data)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}

	fmt.Println(pp.String())

	data, err = ioutil.ReadFile("../RE/Packets/MusicScore/1.08.0/PUMP IT UP: PRIME_226.bin")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	pp = ScoreBoardPacket2{}
	err = pp.FromBinary(data)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}

	fmt.Println(pp.String())
}
