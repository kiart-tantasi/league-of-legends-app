package match

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/pkg/api"
	"io"
	"net/http"
	"strings"
	"sync"
)

type MatchResponse struct {
	Puuid string `json:"puuid"`
}

type MatchDetailResponse struct {
	Info MatchInfo `json:"info"`
}
type MatchInfo struct {
	Participants []Participant `json:"participants"`
	GameMode     string        `json:"gameMode"`
	GameCreation int           `json:"gameCreation"`
}
type Participant struct {
	RiotIdGameName string `json:"riotIdGameName"`
	RiotIdtagLine  string `json:"riotIdTagLine"`
	Kills          int    `json:"kills"`
	Deaths         int    `json:"deaths"`
	Assists        int    `json:"assists"`
	ChampionName   string `json:"championName"`
	Win            bool   `json:"win"`
	Item0          int    `json:"item0"`
	Item1          int    `json:"item1"`
	Item2          int    `json:"item2"`
	Item3          int    `json:"item3"`
	Item4          int    `json:"item4"`
	Item5          int    `json:"item5"`
	Item6          int    `json:"item6"`
}

func getMatches(gameName, tagLine string) (string, error) {
	puuid, err := getPuuid(gameName, tagLine)
	if err != nil {
		return "", err
	}
	matchIds, err := getMatchIds(puuid)
	if err != nil {
		return "", err
	}
	matches, err := getMatchDetails(matchIds)
	if err != nil {
		return "", nil
	}
	return strings.Join(matches, "-"), nil
}

func getPuuid(gameName, tagLine string) (string, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", getRitoRegionAccount(), gameName, tagLine)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Riot-Token", getRiotApiKey())
	res, err := (api.NewHttpClient()).Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var matchResponse MatchResponse
	err = json.Unmarshal(bytes, &matchResponse)
	if err != nil {
		return "", err
	}
	return matchResponse.Puuid, nil
}

func getMatchIds(puuid string) (*[]string, error) {
	url := fmt.Sprintf(
		"https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d", getRiotRegionMatch(), puuid, getRiotMatchAmount())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Riot-Token", getRiotApiKey())
	res, err := (api.NewHttpClient()).Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var matchIds []string
	json.Unmarshal(bytes, &matchIds)
	return &matchIds, nil
}

func getMatchDetails(matchIds *[]string) ([]string, error) {
	responses := make([]*MatchDetailResponse, len(*matchIds))
	limitChannel := make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(len(*matchIds))
	for i, matchId := range *matchIds {
		limitChannel <- 0
		go func(matchId string, responses []*MatchDetailResponse, i int) {
			defer wg.Done()
			err := getMatchDetail(matchId, responses, i)
			if err != nil {
				fmt.Println("getMatchDetail error:", err)
			}
			<-limitChannel
		}(matchId, responses, i)
	}
	wg.Wait()

	// TODO: remodel

	return nil, errors.New("not implemented")
}

func getMatchDetail(matchId string, responses []*MatchDetailResponse, index int) error {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s", getRiotRegionMatch(), matchId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Riot-Token", getRiotApiKey())
	res, err := (api.NewHttpClient()).Do(req)
	// why check error before defer: https://stackoverflow.com/a/16280362/21331113
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var matchDetailResponse MatchDetailResponse
	err = json.Unmarshal(bytes, &matchDetailResponse)
	if err != nil {
		return err
	}
	responses[index] = &matchDetailResponse
	return nil
}
