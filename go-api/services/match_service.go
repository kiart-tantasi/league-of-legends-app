package services

import (
	"encoding/json"
	"fmt"
	"go-api/configs"
	"io"
	"net/http"
	"strings"
	"time"
)

type MatchResponse struct {
	Puuid string `json:"puuid"`
}

type MatchService struct{}

func (matchService *MatchService) GetMatches(gameName, tagLine string) (string, error) {
	puuid, err := getPuuid(gameName, tagLine)
	if err != nil {
		return "", err
	}
	fmt.Println("puuid:", puuid)
	matchIds := getMatchIds(puuid)
	matches := getMatches(matchIds)
	return strings.Join(matches, "-"), nil
}

func getPuuid(gameName, tagLine string) (string, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", configs.GetRegionAccount(), gameName, tagLine)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Riot-Token", configs.GetRiotApiKey())
	res, err := (getHttpClient()).Do(req)
	if err != nil {
		return "", err
	}
	jsonBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var matchResponse MatchResponse
	err = json.Unmarshal(jsonBytes, &matchResponse)
	if err != nil {
		return "", err
	}
	return matchResponse.Puuid, nil
}

func getMatchIds(puuid string) []string {
	return []string{"id1", "id2", "id3"}
}

func getMatches(matchIds []string) []string {
	return matchIds
}

func getHttpClient() *http.Client {
	client := http.Client{
		Timeout: time.Duration(5000 * time.Millisecond),
	}
	return &client
}
