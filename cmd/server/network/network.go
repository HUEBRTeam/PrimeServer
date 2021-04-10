package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/proto"
	"github.com/ajg/form"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

func RetrieveProfile(apikey string, accesscode string, address string, pm *ProfileManager.ProfileManager) (proto.ProfilePacket, error) {
	u, err := url.Parse(address)
	if err != nil {
		return proto.ProfilePacket{}, err
	}

	u.Path = path.Join(u.Path, "getprofile")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return proto.ProfilePacket{}, err
	}

	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("access_code", accesscode)
	req.URL.RawQuery = query.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return proto.ProfilePacket{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return proto.ProfilePacket{}, err
	}

	profpacket, err := pm.GetStorageBackend().GetProfile(accesscode)
	if err != nil {
		profpacket = *proto.MakeProfilePacketDefault("", accesscode)
		err = nil
	}

	err = json.Unmarshal(body, &profpacket) // may have to switch profpacket.AccessCode
	return profpacket, err
}

func RetrieveWorldBest(apikey string, address string, scoretype string) (*proto.WorldBestPacket, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "getworldbest")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("scoretype", scoretype)
	req.URL.RawQuery = query.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	wbpacket := proto.MakeWorldBestPacket([]proto.WorldBestScore{})
	err = json.Unmarshal(body, &wbpacket)
	return wbpacket, err
}

func RetrieveRankMode(apikey string, address string, scoretype string) (*proto.RankModePacket, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "getrankmode")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("api_key", apikey)
	query.Add("scoretype", scoretype)
	req.URL.RawQuery = query.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rnkpacket := proto.MakeRankModePacket([]proto.SongRank{})
	err = json.Unmarshal(body, &rnkpacket)

	return rnkpacket, err
}

func SubmitScore(apikey string, address string, score proto.ScoreBoardPacket2, accesscode string) error {
	u, err := url.Parse(address)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "submit")
	values, err := form.EncodeToValues(score)
	if err != nil {
		return err
	}
	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	values.Set("AccessCode", accesscode)

	resp, err := http.PostForm(u.String(), values)
	//fmt.Println("Score values: %+v", values)
	if err != nil {
		return err
	}

	fmt.Println("Score response:")
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))
	resp.Body.Close()

	return nil
}

func SubmitProfile(apikey string, address string, profile proto.ProfilePacket, accesscode string) error {
	u, err := url.Parse(address)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "saveprofile")
	values, err := form.EncodeToValues(profile)

	if err != nil {
		return err
	}

	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	values.Set("AccessCode", accesscode)

	// remove this after debugging
	f, err := os.Create("submittedprofile.log")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(fmt.Sprintf("%+v\n", values))
	w.Flush()

	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return err
	}
	fmt.Println("Profile response:")
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
	return nil
}

func SubmitMachineInfo(apikey string, address string, profile proto.MachineInfoPacket) error {
	u, err := url.Parse(address)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "submitmachineinfo")
	values, err := form.EncodeToValues(profile)

	if err != nil {
		return err
	}

	// for any extras just do values.Set(key, value)
	values.Set("api_key", apikey)
	resp, err := http.PostForm(u.String(), values)
	if err != nil {
		return err
	}
	//fmt.Println("Machine Info response:")
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	resp.Body.Close()

	return nil
}
