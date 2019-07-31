package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/proto"
)

func RetrieveProfile(apikey string, accesscode string, address string, pm *ProfileManager.ProfileManager) (profpacket proto.ProfilePacket, err error) {
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
	profpacket, err = pm.GetStorageBackend().GetProfile(accesscode)
	if err != nil {
		return
	}
	_ = json.Unmarshal(body, &profpacket) // may have to switch profpacket.AccessCode
	return
}

func RetrieveWorldBest(apikey string, address string, scoretype string) (wbpacket *proto.WorldBestPacket, err error) {
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
	wbpacket = proto.MakeWorldBestPacket([]proto.WorldBestScore{})
	_ = json.Unmarshal(body, &wbpacket)
	return
}

func RetrieveRankMode(apikey string, address string, scoretype string) (rnkpacket *proto.RankModePacket, err error) {
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
	wbpacket := proto.MakeRankModePacket([]proto.SongRank{})
	_ = json.Unmarshal(body, &wbpacket)
	return
}

func SubmitScore(apikey string, address string, score proto.ScoreBoardPacket2, accesscode string) (err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "submit")
	values := url.Values{}
	val := reflect.ValueOf(score).Elem()
	t := val.Type()
	for i := 0; i < val.NumField(); i++ { // iterate through struct fields and convert everything to strings
		values.Set(t.Field(i).Name, fmt.Sprint(val.Field(i)))
	}
	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	values.Set("AccessCode", accesscode)
	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func SubmitProfile(apikey string, address string, profile proto.ProfilePacket, accesscode string) (err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "saveprofile")
	values := url.Values{}
	val := reflect.ValueOf(profile).Elem()
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		values.Set(t.Field(i).Name, fmt.Sprint(val.Field(i)))
	}
	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	values.Set("AccessCode", accesscode)
	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func SubmitMachineInfo(apikey string, address string, profile proto.MachineInfoPacket) (err error) {
	u, err := url.Parse(address)
	if err != nil {
		return
	}
	u.Path = path.Join(u.Path, "submitmachineinfo")
	values := url.Values{}
	val := reflect.ValueOf(profile).Elem()
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		values.Set(t.Field(i).Name, fmt.Sprint(val.Field(i)))
	}
	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}
