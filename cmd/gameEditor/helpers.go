package main

const (
	primeServerAddress = "prime_daemon.piugame.com"
	piugameAddress     = "www2.piugame.com"
	versionStart       = "V1."

	maxStringLength = 2048
)

func CompareByteArray(a, b []byte) bool {
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func SearchGadget(data []byte, feature []byte) (offset int) {
	total := len(data)
	fTotal := len(feature)

	total -= fTotal

	for i := 0; i < total; i++ {
		if CompareByteArray(data[i:i+fTotal], feature) {
			return i
		}
	}

	return -1
}

func SearchAllGadgets(data []byte, feature []byte) []int {
	total := len(data)
	fTotal := len(feature)
	total -= fTotal

	gadgets := make([]int, 0)

	for i := 0; i < total; i++ {
		if CompareByteArray(data[i:i+fTotal], feature) {
			gadgets = append(gadgets, i)
		}
	}

	return gadgets
}

func GetStringAt(data []byte, offset int) string {
	s := ""
	for i := offset; i < offset+maxStringLength; i++ {
		if data[i] == 0x00 {
			break
		}
		s += string(data[i])
	}

	return s
}

func SetStringAt(data []byte, offset int, s string) {
	for i, v := range s {
		data[offset+i] = byte(v)
	}

	data[offset+len(s)] = 0x00
}
