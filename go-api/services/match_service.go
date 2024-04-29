package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type MatchService struct{}

func (matchService *MatchService) GetMatches(gameName, tagLine string) string {
	puuid := getPuuid(gameName, tagLine)
	matchIds := getMatchIds(puuid)
	matches := getMatches(matchIds)
	return strings.Join(matches, "-")
}

func getPuuid(gameName, tagLine string) string {
	// TODO: handle config
	var riotConfig *RiotConfig
	url := fmt.Sprintf("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", riotConfig.getRegionAccount(), gameName, tagLine)
	// TODO: add header "X-Riot-Token"
	if res, err := http.Get(url); err != nil {
		panic(err)
	} else {
		fmt.Println(res.Body)
		// TODO: parse json
	}
	return ""
}

func getMatchIds(puuid string) []string {
	return []string{"id1", "id2", "id3"}
}

func getMatches(matchIds []string) []string {
	for matchId := range matchIds {
		fmt.Println("matchId:", matchId)
	}
	return matchIds
}

// riot config
type RiotConfig struct{}

func (*RiotConfig) getRiotApiKey() string {
	return os.Getenv("RIOT_API_KEY")
}
func (*RiotConfig) getRegionAccount() string {
	return os.Getenv("RIOT_API_REGION_ACCOUNT")
}
func (*RiotConfig) getRegionMatch() string {
	return os.Getenv("RIOT_API_REGION_MATCH")
}
func (*RiotConfig) getMatchAmount() int {
	if matchAmount, err := strconv.Atoi(os.Getenv("RIOT_MATCH_AMOUNT")); err != nil {
		panic("RIOT_MATCH_AMOUNT is not set properly")
	} else {
		return matchAmount
	}
}
