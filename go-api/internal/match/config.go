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
func GetRiotMatchAmount() (int, error) {
	if matchAmount, err := strconv.Atoi(env.GetEnv("RIOT_MATCH_AMOUNT", "20")); err != nil {
		return 0, err
	} else {
		return matchAmount, nil
	}
}
