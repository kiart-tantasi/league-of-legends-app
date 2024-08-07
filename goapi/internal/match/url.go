package match

import (
	"fmt"
	"os"
)

// different between actual api urls and mockapi urls is that
// actual api urls contain sub-domain (for region option) while mockapi urls do not

func getRiotAccountApiUrl(gameName, tagLine string) string {
	defaultUrl := "http://localhost:8090/riot/account/v1/accounts/by-riot-id/%s/%s"
	url := os.Getenv("RIOT_ACCOUNT_API_URL")
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotAccountRegion(), gameName, tagLine)
	}
	return fmt.Sprintf(defaultUrl, gameName, tagLine)
}

func getRiotMatchIdsApiUrl(puuid string) string {
	defaultUrl := "http://localhost:8090/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d"
	url := os.Getenv("RIOT_MATCH_IDS_API_URL")
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotMatchRegion(), puuid, getRiotMatchAmount())
	}
	return fmt.Sprintf(defaultUrl, puuid, getRiotMatchAmount())
}

func getRiotMatchDetailApiUrl(matchId string) string {
	defaultUrl := "http://localhost:8090/lol/match/v5/matches/%s"
	url := os.Getenv("RIOT_MATCH_DETAIL_API_URL")
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotMatchRegion(), matchId)
	}
	return fmt.Sprintf(defaultUrl, matchId)
}
