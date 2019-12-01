package proto

import (
	"bytes"
	"encoding/json"
	//"fmt"
)

type GenericPacket interface {
	FromBinary([]byte) error
	ToBinary() []byte
	GetType() uint32
	GetName() string
}

type PIUString12 [12]uint8
type PIUString16 [16]uint8
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

func MakePIUString16(data string) PIUString16 {
	p := PIUString16{}
	p2 := makePIUString(data, 16)
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

// region JSON Handlers

func (n PIUString12) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUString12) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 12)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

func (n PIUNickname) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUNickname) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 12)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

func (n PIUString16) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUString16) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 16)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

func (n PIUMacAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUMacAddress) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 20)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

func (n PIUString32) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUString32) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 32)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

func (n PIUString128) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *PIUString128) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	//fmt.Printf("String: %v\n", s)
	a := makePIUString(s, 128)
	copy(n[:], a)
	//fmt.Printf("PIUString: %v\n", n.String())
	return nil
}

// endregion

// region To String Handlers
func (n PIUString12) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 12
	}
	return string(n[:s])
}

func (n PIUNickname) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 12
	}
	return string(n[:s])
}

func (n PIUString16) String() string {
	s := bytes.Index(n[:], []byte{0})
	if s == -1 {
		s = 16
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
		s = 32
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

// region Form Encoded Marshal / Unmarshal
func (n PIUString12) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *PIUString12) UnmarshalText(text []byte) error {
	a := makePIUString(string(text), 12)
	copy(n[:], a)

	return nil
}

func (n PIUNickname) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *PIUNickname) UnmarshalText(text []byte) error {
	a := makePIUString(string(text), 12)
	copy(n[:], a)

	return nil
}
func (n PIUString16) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *PIUString16) UnmarshalText(text []byte) error {
	a := makePIUString(string(text), 16)
	copy(n[:], a)

	return nil
}

func (n PIUString32) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *PIUString32) UnmarshalText(text []byte) error {
	a := makePIUString(string(text), 32)
	copy(n[:], a)

	return nil
}

func (n PIUString128) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *PIUString128) UnmarshalText(text []byte) error {
	a := makePIUString(string(text), 128)
	copy(n[:], a)

	return nil
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
