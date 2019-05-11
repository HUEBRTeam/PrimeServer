package proto

import "bytes"

type PIUString12 [12]uint8
type PIUMacAddress [20]uint8
type PIUString32 [32]uint8
type PIUString128 [128]uint8
type PIUNickname PIUString12

func MakePIUString12(data string) PIUString12 {
	p := PIUString12{}
	p2 := makePIUString(data, 12)
	copy(p[:], p2)
	return p
}

func MakePIUString32(data string) PIUString32 {
	p := PIUString32{}
	p2 := makePIUString(data, 32)
	copy(p[:], p2)
	return p
}

func MakePIUString128(data string) PIUString128 {
	p := PIUString128{}
	p2 := makePIUString(data, 128)
	copy(p[:], p2)
	return p
}

func MakePIUMacAddress(data string) PIUMacAddress {
	p := PIUMacAddress{}
	p2 := makePIUString(data, 20)
	copy(p[:], p2)
	return p
}

func MakePIUNickName(nickname string) PIUNickname {
	return PIUNickname(MakePIUString12(nickname))
}

// region To String Handlers
func (n PIUString12) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 12
	}
	return string(n[:s])
}

func (n PIUMacAddress) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 20
	}
	return string(n[:s])
}

func (n PIUString32) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 20
	}
	return string(n[:s])
}

func (n PIUString128) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 128
	}
	return string(n[:s])
}

// endregion

// region Private Makers
func makePIUString(data string, length int) []byte {
	p := make([]byte, length)
	for i := range p {
		if len(data) > i {
			p[i] = data[i]
		} else {
			p[i] = 0x00
		}
	}

	return p
}

// endregion
