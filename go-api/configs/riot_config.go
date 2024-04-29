package configs

import (
	"os"
	"strconv"
)

type RiotConfig struct{}

func (*RiotConfig) GetRiotApiKey() string {
	return os.Getenv("RIOT_API_KEY")
}
func (*RiotConfig) GetRegionAccount() string {
	return os.Getenv("RIOT_API_REGION_ACCOUNT")
}
func (*RiotConfig) GetRegionMatch() string {
	return os.Getenv("RIOT_API_REGION_MATCH")
}
func (*RiotConfig) GetMatchAmount() (int, error) {
	if matchAmount, err := strconv.Atoi(os.Getenv("RIOT_MATCH_AMOUNT")); err != nil {
		return 0, err
	} else {
		return matchAmount, nil
	}
}
