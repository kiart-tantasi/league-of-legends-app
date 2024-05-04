package match

import (
	"fmt"
	"go-api/pkg/env"
)

// different between actual api url and mock-api's url is that
// actual api url contains sub-domain (for region option) while mock-api url does not

func newRiotAccountApiUrl(gameName, tagLine string) string {
	defaultUrl := "http://localhost:8090/riot/account/v1/accounts/by-riot-id/%s/%s"
	url := env.GetEnv("RIOT_ACCOUNT_API_URL", defaultUrl)
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotRegionAccount(), gameName, tagLine)
	}
	return fmt.Sprintf(defaultUrl, gameName, tagLine)
}

func newRiotMatchIdsApiUrl(puuid string) string {
	defaultUrl := "http://localhost:8090/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d"
	url := env.GetEnv("RIOT_MATCH_IDS_API_URL", defaultUrl)
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotRegionMatch(), puuid, getRiotMatchAmount())
	}
	return fmt.Sprintf(defaultUrl, puuid, getRiotMatchAmount())
}

func newRiotMatchDetailApiUrl(matchId string) string {
	defaultUrl := "http://localhost:8090/lol/match/v5/matches/%s"
	url := env.GetEnv("RIOT_MATCH_DETAIL_API_URL", defaultUrl)
	if url != defaultUrl {
		return fmt.Sprintf(url, getRiotRegionMatch(), matchId)
	}
	return fmt.Sprintf(defaultUrl, matchId)
}
