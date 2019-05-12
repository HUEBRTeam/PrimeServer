package tests

import (
	"github.com/HUEBRTeam/PrimeServer/proto"
	"golang.org/x/crypto/nacl/box"
	"io/ioutil"
	"testing"
)

func TestSodium(t *testing.T) {
	unenc, _ := ioutil.ReadFile("../RE/Packets/RawSample/Unencrypted.bin")
	nonce, _ := ioutil.ReadFile("../RE/Packets/RawSample/Nounce.bin")
	enc, _ := ioutil.ReadFile("../RE/Packets/RawSample/Encrypted.bin")

	var nonceBytes [24]byte

	copy(nonceBytes[:], nonce)

	outEnc := box.Seal(nil, unenc, &nonceBytes, &proto.ServerPublicKeyBytes, &proto.ServerPrivateKeyBytes)

	if len(outEnc) != len(enc) {
		t.Fatalf("Expected output length to be %d got %d", len(enc), len(outEnc))
	}

	for i, v := range outEnc {
		if enc[i] != v {
			t.Fatalf("Encryption Error. Byte at position %d is wrong.", i)
		}
	}

	t.Log("Encryption OK")

	outDec, ok := box.Open(nil, enc, &nonceBytes, &proto.ServerPublicKeyBytes, &proto.ServerPrivateKeyBytes)

	if !ok {
		t.Fatalf("Decryption error")
	}

	for i, v := range outDec {
		if unenc[i] != v {
			t.Fatalf("Decryption Error. Byte at position %d is wrong.", i)
		}
	}
	t.Log("Decryption OK")
}
