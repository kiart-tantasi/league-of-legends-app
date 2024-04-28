package services

import (
	"fmt"
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
