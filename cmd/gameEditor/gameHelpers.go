package main

import "fmt"

func GetGameVersion(data []byte) (string, error) {
	s := SearchAllGadgets(data, []byte(versionStart))

	if len(s) == 0 {
		return "", fmt.Errorf("game version not found")
	}

	for _, v := range s {
		o := GetStringAt(data, v)
		if len(o) == 7 {
			return o, nil
		}
	}

	return "", fmt.Errorf("game version not found")
}

func FindPrimeDaemonURL(data []byte) int {
	return SearchGadget(data, []byte(primeServerAddress))
}

func FindPrimeServerIP(data []byte) (int, string) {
	i := FindPrimeDaemonURL(data)
	i += len(primeServerAddress) + 1
	return i, GetStringAt(data, i)
}

func SetPrimeServerIP(data []byte, ip string) bool {
	if len(ip) > 14 {
		return false
	}

	idx, _ := FindPrimeServerIP(data)

	SetStringAt(data, idx, ip)

	return true
}
