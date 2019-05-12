package proto

import (
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
