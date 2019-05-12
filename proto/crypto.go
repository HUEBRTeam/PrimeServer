package proto

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/quan-to/slog"
	"golang.org/x/crypto/nacl/box"
)

const ServerPrivateKey = "\x93\xb2\xf8\xa0\xe0\x95\x49\xdf\xaf\xfe\x36\x56\xe6\xad\x9c\x9a\x53\x9b\xf6\x63\xd3\x7c\x08\xf0\xa4\xe6\x29\x3c\xbf\xfa\x64\x56"
const ServerPublicKey = "\xb0\xeb\x81\x47\x51\x0a\x2e\x20\x32\xed\x2f\xc0\xf9\xd4\xeb\x88\xb0\xca\x2d\xbd\xe7\xd6\x4e\xe2\xd0\x83\x4c\x3b\xb6\xfb\x31\x55"

var ServerPrivateKeyBytes [32]byte
var ServerPublicKeyBytes [32]byte
var clog = slog.Scope("Crypto")

func init() {
	copy(ServerPrivateKeyBytes[:], ServerPrivateKey)
	copy(ServerPublicKeyBytes[:], ServerPublicKey)
}

func DecryptPacket(packet []byte) ([]byte, bool) {
	var nonceBytes [24]byte

	nonce := packet[:24]
	data := packet[24:]

	copy(nonceBytes[:], nonce)

	return box.Open(nil, data, &nonceBytes, &ServerPublicKeyBytes, &ServerPrivateKeyBytes)
}

func EncryptPacket(packet []byte) []byte {
	var nonceBytes [24]byte

	_, err := rand.Read(nonceBytes[:])

	if err != nil {
		clog.Error("Error reading from crypto/random: %s", err)
		return nil
	}

	outEnc := box.Seal(nil, packet, &nonceBytes, &ServerPublicKeyBytes, &ServerPrivateKeyBytes)

	outputData := make([]byte, len(outEnc)+24+4)

	binary.LittleEndian.PutUint32(outputData, uint32(len(outputData)))

	copy(outputData[4:], nonceBytes[:])
	copy(outputData[4+24:], outEnc)

	return outputData
}
