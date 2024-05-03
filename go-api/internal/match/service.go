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
	if res.StatusCode != 200 {
		return "", fmt.Errorf("puuid response status code is not 200 for %s #%s", gameName, tagLine)
	}
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
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("match ids response status code is not 200 for %s", puuid)
	}
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

func getMatchesResponse(matchIds *[]string, puuid string) (*MatchesResponseV1, error) {
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
				fmt.Printf("getMatchDetail error for match id %s: %s\n", matchId, err)
			} else {
				responses[i] = response
			}
			<-limitChannel
		}(matchId, responses, i)
	}
	wg.Wait()
	return mapToResponse(responses, puuid), nil
}

func getMatchDetail(matchId string) (*RiotMatchDetailResponse, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s", getRiotRegionMatch(), matchId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Riot-Token", getRiotApiKey())
	res, err := (api.NewHttpClient()).Do(req)
	// why check error before `defer res.Body.Close()`: https://stackoverflow.com/a/16280362/21331113
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("match detail response status code is not 200")
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

func mapToResponse(responses []*RiotMatchDetailResponse, puuid string) *MatchesResponseV1 {
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
	return &MatchesResponseV1{MatchDetailList: list}
}
