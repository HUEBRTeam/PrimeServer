package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/proto"
)

func RetrieveProfile(apikey string, accesscode string, address string, pm ProfileManager.ProfileManager) (profpacket proto.ProfilePacket, err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "getprofile")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("access_code", accesscode)
	req.URL.RawQuery = query.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	profpacket := pm.GetProfile(accesscode)
	_ = json.Unmarshal(body, &profpacket)
	return
}

func RetrieveWorldBest(apikey string, address string, scoretype string) (wbpacket proto.WorldBestPacket, err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "getworldbest")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("scoretype", scoretype)
	req.URL.RawQuery = query.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	wbpacket := proto.MakeWorldBestPacket([]WorldBestScore{})
	_ = json.Unmarshal(body, &wbpacket)
	return
}

func RetrieveRankMode(apikey string, address string, scoretype string) (rnkpacket proto.RankModePacket, err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "getrankmode")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("scoretype", scoretype)
	req.URL.RawQuery = query.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	wbpacket := proto.MakeRankModePacket([]SongRank{})
	_ = json.Unmarshal(body, &wbpacket)
	return
}

func SubmitScore(apikey string, address string, score proto.ScoreBoardPacket2) (err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "submit")
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("access_code", accesscode)
	req.URL.RawQuery = query.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	profpacket := pm.GetProfile(accesscode)
	_ = json.Unmarshal(body, &profpacket)
	return
}

func SubmitProfile(apikey string, address string, profile proto.ProfilePacket) (err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "saveprofile")
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return
	}
	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("access_code", accesscode)
	req.URL.RawQuery = query.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	profpacket := pm.GetProfile(accesscode)
	_ = json.Unmarshal(body, &profpacket)
	return
}