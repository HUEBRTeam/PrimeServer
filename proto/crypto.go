package proto

import (
	"encoding/binary"
	"github.com/HUEBRTeam/PrimeServer"
	"golang.org/x/crypto/nacl/box"
)

var LastNounce []byte

func DecryptPacket(packet []byte) ([]byte, bool) {
	var nonceBytes [24]byte

	nonce := packet[:24]
	data := packet[24:]
	LastNounce = nonce

	copy(nonceBytes[:], nonce)

	return box.Open(nil, data, &nonceBytes, &PrimeServer.ServerPublicKeyBytes, &PrimeServer.ServerPrivateKeyBytes)
}

func EncryptPacket(packet []byte) []byte {
	var nonceBytes [24]byte
	copy(nonceBytes[:], LastNounce)

	outEnc := box.Seal(nil, packet, &nonceBytes, &PrimeServer.ServerPublicKeyBytes, &PrimeServer.ServerPrivateKeyBytes)

	outputData := make([]byte, len(outEnc)+24+4)

	binary.LittleEndian.PutUint32(outputData, uint32(len(outputData)))

	copy(outputData[4:], nonceBytes[:])
	copy(outputData[4+24:], outEnc)

	return outputData
}
