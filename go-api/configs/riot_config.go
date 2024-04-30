package configs

import (
	"go-api/utils"
	"strconv"
)

func GetRiotApiKey() string {
	return utils.GetEnv("RIOT_API_KEY", "please retrieve api key from https://developer.riotgames.com/")
}
func GetRegionAccount() string {
	return utils.GetEnv("RIOT_API_REGION_ACCOUNT", "asia")
}
func GetRegionMatch() string {
	return utils.GetEnv("RIOT_API_REGION_MATCH", "sea")
}
func GetMatchAmount() (int, error) {
	if matchAmount, err := strconv.Atoi(utils.GetEnv("RIOT_MATCH_AMOUNT", "20")); err != nil {
		return 0, err
	} else {
		return matchAmount, nil
	}
}
