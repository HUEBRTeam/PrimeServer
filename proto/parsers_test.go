package proto

import (
	"io/ioutil"
	"testing"
)

func TestMachineInfoPacket_FromBinary(t *testing.T) {
	data, _ := ioutil.ReadFile("../RE/Packets/Start/sodium_53_out_pkt.bin")

	mip := MachineInfoPacket{}
	err := mip.FromBinary(data)

	if err != nil {
		t.Fatalf("Error parsing packet: %s", err)
	}

	// TODO: Test field values
}
