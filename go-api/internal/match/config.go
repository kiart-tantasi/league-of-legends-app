package match

import (
	"os"
	"strconv"
)

func getRiotApiKey() string {
	return os.Getenv("RIOT_API_KEY")
}
func getRiotAccountRegion() string {
	return os.Getenv("RIOT_API_REGION_ACCOUNT")
}
func getRiotMatchRegion() string {
	return os.Getenv("RIOT_API_REGION_MATCH")
}
func getRiotMatchAmount() int {
	matchAmount, err := strconv.Atoi(os.Getenv("RIOT_MATCH_AMOUNT"))
	if err != nil {
		panic(err)
	}
	return matchAmount
}
