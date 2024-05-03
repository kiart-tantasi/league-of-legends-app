package match

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/pkg/api"
	"io"
	"net/http"
	"sync"
)

type RiotMatchResponse struct {
	Puuid string `json:"puuid"`
}

type RiotMatchDetailResponse struct {
	Info RiotMatchInfo `json:"info"`
}
type RiotMatchInfo struct {
	Participants []RiotParticipant `json:"participants"`
	GameMode     string            `json:"gameMode"`
	GameCreation int               `json:"gameCreation"`
}

type RiotParticipant struct {
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
	Puuid          string `json:"puuid"`
}

func (riotParticipant *RiotParticipant) getItemIds() *[]int {
	itemIds := []int{}
	if riotParticipant.Item0 != 0 {
		itemIds = append(itemIds, riotParticipant.Item0)
	}
	if riotParticipant.Item1 != 0 {
		itemIds = append(itemIds, riotParticipant.Item1)
	}
	if riotParticipant.Item2 != 0 {
		itemIds = append(itemIds, riotParticipant.Item2)
	}
	if riotParticipant.Item3 != 0 {
		itemIds = append(itemIds, riotParticipant.Item3)
	}
	if riotParticipant.Item4 != 0 {
		itemIds = append(itemIds, riotParticipant.Item4)
	}
	if riotParticipant.Item5 != 0 {
		itemIds = append(itemIds, riotParticipant.Item5)
	}
	if riotParticipant.Item6 != 0 {
		itemIds = append(itemIds, riotParticipant.Item6)
	}
	return &itemIds
}

type MatcesResponseV1 struct {
	MatchDetailList []MatchDetailV1 `json:"matchDetailList"`
}
type MatchDetailV1 struct {
	ChampionName    string           `json:"championName"`
	Kills           int              `json:"kills"`
	Deaths          int              `json:"deaths"`
	Assists         int              `json:"assists"`
	Win             bool             `json:"win"`
	GameMode        string           `json:"gameMode"`
	GameCreation    int              `json:"gameCreation"`
	ParticipantList *[]ParticipantV1 `json:"participantList"`
	ItemIds         []int            `json:"itemIds"`
}

type ParticipantV1 struct {
	GameName     string `json:"gameName"`
	TagLine      string `json:"tagLine"`
	ChampionName string `json:"championName"`
	Kills        int    `json:"kills"`
	Deaths       int    `json:"deaths"`
	Assists      int    `json:"assists"`
	Win          bool   `json:"win"`
	ItemIds      []int  `json:"itemIds"`
}

func getMatchesV1(gameName, tagLine string) ([]byte, error) {
	puuid, err := getPuuid(gameName, tagLine)
	if err != nil {
		return nil, err
	}
	matchIds, err := getMatchIds(puuid)
	if err != nil {
		return nil, err
	}
	matchesResponse, err := getMatchesResponse(matchIds, puuid)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(&matchesResponse)
	if err != nil {
		return nil, err
	}
	return bytes, nil
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
	var matchResponse RiotMatchResponse
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
	err = json.Unmarshal(bytes, &matchIds)
	if err != nil {
		return nil, err
	}
	return &matchIds, nil
}

func getMatchesResponse(matchIds *[]string, puuid string) (*MatcesResponseV1, error) {
	responses := make([]*RiotMatchDetailResponse, len(*matchIds))
	limitChannel := make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(len(*matchIds))
	for i, matchId := range *matchIds {
		limitChannel <- 0
		go func(matchId string, responses []*RiotMatchDetailResponse, i int) {
			defer wg.Done()
			response, err := getMatchDetail(matchId)
			if err != nil {
				fmt.Println("getMatchDetail error:", err)
			} else {
				responses[i] = response
			}
			<-limitChannel
		}(matchId, responses, i)
	}
	wg.Wait()
	// ==================
	// TODO: check default value of slice (json-decoded)
	// TODO: create a separate func
	list := []MatchDetailV1{}
	for _, response := range responses {
		if response == nil {
			continue
		}

		matchDetail := &MatchDetailV1{}
		participants := []ParticipantV1{}
		for _, parti := range response.Info.Participants {
			// all cases
			participant := ParticipantV1{
				GameName:     parti.RiotIdGameName,
				TagLine:      parti.RiotIdtagLine,
				ChampionName: parti.ChampionName,
				Kills:        parti.Kills,
				Assists:      parti.Assists,
				Deaths:       parti.Deaths,
				Win:          parti.Win,
				ItemIds:      *parti.getItemIds(),
			}
			participants = append(participants, participant)
			// id owner case
			if parti.Puuid == puuid {
				matchDetail.ChampionName = parti.ChampionName
				matchDetail.Kills = parti.Kills
				matchDetail.Assists = parti.Assists
				matchDetail.Deaths = parti.Deaths
				matchDetail.Win = parti.Win
				matchDetail.GameMode = response.Info.GameMode
				matchDetail.GameCreation = response.Info.GameCreation
				matchDetail.ItemIds = *parti.getItemIds()
			}
		}
		matchDetail.ParticipantList = &participants
		list = append(list, *matchDetail)
	}
	// ==================
	return &MatcesResponseV1{MatchDetailList: list}, nil
}

func getMatchDetail(matchId string) (*RiotMatchDetailResponse, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s", getRiotRegionMatch(), matchId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Riot-Token", getRiotApiKey())
	res, err := (api.NewHttpClient()).Do(req)
	// why check error before defer: https://stackoverflow.com/a/16280362/21331113
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("response status code is not 200")
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var matchDetailResponse RiotMatchDetailResponse
	err = json.Unmarshal(bytes, &matchDetailResponse)
	if err != nil {
		return nil, err
	}
	return &matchDetailResponse, nil
}
