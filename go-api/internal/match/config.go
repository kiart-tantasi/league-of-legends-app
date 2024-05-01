package match

import (
	"go-api/pkg/env"
	"strconv"
)

func GetRiotApiKey() string {
	return env.GetEnv("RIOT_API_KEY", "please retrieve api key from https://developer.riotgames.com/")
}
func GetRitoRegionAccount() string {
	return env.GetEnv("RIOT_API_REGION_ACCOUNT", "asia")
}
func GetRiotRegionMatch() string {
	return env.GetEnv("RIOT_API_REGION_MATCH", "sea")
}
func GetRiotMatchAmount() int {
	// when RIOT_MATCH_AMOUNT is set correctly, we can assume there is no error
	matchAmount, _ := strconv.Atoi(env.GetEnv("RIOT_MATCH_AMOUNT", "5"))
	return matchAmount
}
