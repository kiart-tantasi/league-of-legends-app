package match

type RiotMatchResponse struct {
	Puuid string `json:"puuid"`
}

type RiotMatchDetailResponse struct {
	Info RiotMatchInfo `json:"info"`
}
type RiotMatchInfo struct {
	Participants []RiotParticipant `json:"participants"`
	GameMode     string            `json:"gameMode"`
	GameCreation int               `json:"gameCreation"`
}

type RiotParticipant struct {
	RiotIdGameName string `json:"riotIdGameName"`
	RiotIdtagLine  string `json:"riotIdTagLine"`
	Kills          int    `json:"kills"`
	Deaths         int    `json:"deaths"`
	Assists        int    `json:"assists"`
	ChampionName   string `json:"championName"`
	Win            bool   `json:"win"`
	Item0          int    `json:"item0"`
	Item1          int    `json:"item1"`
	Item2          int    `json:"item2"`
	Item3          int    `json:"item3"`
	Item4          int    `json:"item4"`
	Item5          int    `json:"item5"`
	Item6          int    `json:"item6"`
	Puuid          string `json:"puuid"`
	// newly added (feat/participant-v2)
	ChampLevel                     int `json:"champLevel"`
	GoldEarned                     int `json:"goldEarned"`
	DamageDealtToTurrets           int `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int `json:"damageSelfMitigated"`
	MagicDamageDealt               int `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int `json:"magicDamageTaken"`
	PhysicalDamageDealt            int `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int `json:"physicalDamageTaken"`
	TotalDamageDealt               int `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int `json:"totalDamageTaken"`
	TotalHeal                      int `json:"totalHeal"`
	TotalHealsOnTeammates          int `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int `json:"totalMinionsKilled"`
	TrueDamageDealt                int `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int `json:"trueDamageTaken"`
}

type MatchesResponseV1 struct {
	MatchDetailList []MatchDetailV1 `json:"matchDetailList"`
}
type MatchDetailV1 struct {
	ChampionName    string           `json:"championName"`
	Kills           int              `json:"kills"`
	Deaths          int              `json:"deaths"`
	Assists         int              `json:"assists"`
	Win             bool             `json:"win"`
	GameMode        string           `json:"gameMode"`
	GameCreation    int              `json:"gameCreation"`
	ParticipantList *[]ParticipantV2 `json:"participantList"`
	ItemIds         []int            `json:"itemIds"`
}

type ParticipantV1 struct {
	GameName     string `json:"gameName"`
	TagLine      string `json:"tagLine"`
	ChampionName string `json:"championName"`
	Kills        int    `json:"kills"`
	Deaths       int    `json:"deaths"`
	Assists      int    `json:"assists"`
	Win          bool   `json:"win"`
	ItemIds      []int  `json:"itemIds"`
}

type ParticipantV2 struct {
	GameName     string `json:"gameName"`
	TagLine      string `json:"tagLine"`
	ChampionName string `json:"championName"`
	Kills        int    `json:"kills"`
	Deaths       int    `json:"deaths"`
	Assists      int    `json:"assists"`
	Win          bool   `json:"win"`
	ItemIds      []int  `json:"itemIds"`
	// newly added (feat/participant-v2)
	ChampLevel                     int `json:"champLevel"`
	GoldEarned                     int `json:"goldEarned"`
	DamageDealtToTurrets           int `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int `json:"damageSelfMitigated"`
	MagicDamageDealt               int `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int `json:"magicDamageTaken"`
	PhysicalDamageDealt            int `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int `json:"physicalDamageTaken"`
	TotalDamageDealt               int `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int `json:"totalDamageTaken"`
	TotalHeal                      int `json:"totalHeal"`
	TotalHealsOnTeammates          int `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int `json:"totalMinionsKilled"`
	TrueDamageDealt                int `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int `json:"trueDamageTaken"`
}

// sadly, this code will look ugly
func (riotParticipant *RiotParticipant) getItemIds() *[]int {
	itemIds := []int{}
	if riotParticipant.Item0 != 0 {
		itemIds = append(itemIds, riotParticipant.Item0)
	}
	if riotParticipant.Item1 != 0 {
		itemIds = append(itemIds, riotParticipant.Item1)
	}
	if riotParticipant.Item2 != 0 {
		itemIds = append(itemIds, riotParticipant.Item2)
	}
	if riotParticipant.Item3 != 0 {
		itemIds = append(itemIds, riotParticipant.Item3)
	}
	if riotParticipant.Item4 != 0 {
		itemIds = append(itemIds, riotParticipant.Item4)
	}
	if riotParticipant.Item5 != 0 {
		itemIds = append(itemIds, riotParticipant.Item5)
	}
	if riotParticipant.Item6 != 0 {
		itemIds = append(itemIds, riotParticipant.Item6)
	}
	return &itemIds
}
