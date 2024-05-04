package event

type Player struct {
    Tag                          string `json:"tag"`
    Name                         string `json:"name"`
    ExpLevel                     int    `json:"expLevel"`
    Trophies                     int    `json:"trophies"`
    BestTrophies                 int    `json:"bestTrophies"`
    Wins                         int    `json:"wins"`
    Losses                       int    `json:"losses"`
    BattleCount                  int    `json:"battleCount"`
    ThreeCrownWins               int    `json:"threeCrownWins"`
    ChallengeCardsWon            int    `json:"challengeCardsWon"`
    ChallengeMaxWins             int    `json:"challengeMaxWins"`
    TournamentCardsWon           int    `json:"tournamentCardsWon"`
    TournamentBattleCount        int    `json:"tournamentBattleCount"`
    Role                         string `json:"role"`
    Donations                    int    `json:"donations"`
    DonationsReceived            int    `json:"donationsReceived"`
    TotalDonations               int    `json:"totalDonations"`
    WarDayWins                   int    `json:"warDayWins"`
    ClanCardsCollected           int    `json:"clanCardsCollected"`
    Clan                         Clan   `json:"clan"`
    Arena                        Arena  `json:"arena"`
    LeagueStatistics             LeagueStatistics `json:"leagueStatistics"`
    Badges                       []Badge `json:"badges"`
    Achievements                 []Achievement `json:"achievements"`
    Cards                        []Card `json:"cards"`
    SupportCards                 []Card `json:"supportCards"`
    CurrentDeck                  []Card `json:"currentDeck"`
    CurrentDeckSupportCards      []Card `json:"currentDeckSupportCards"`
    CurrentFavouriteCard         Card   `json:"currentFavouriteCard"`
    StarPoints                   int    `json:"starPoints"`
    ExpPoints                    int    `json:"expPoints"`
    LegacyTrophyRoadHighScore    int    `json:"legacyTrophyRoadHighScore"`
    CurrentPathOfLegendSeasonResult  SeasonResult `json:"currentPathOfLegendSeasonResult"`
    LastPathOfLegendSeasonResult     SeasonResult `json:"lastPathOfLegendSeasonResult"`
    BestPathOfLegendSeasonResult     SeasonResult `json:"bestPathOfLegendSeasonResult"`
    TotalExpPoints               int    `json:"totalExpPoints"`
}

type Clan struct {
    Tag     string `json:"tag"`
    Name    string `json:"name"`
    BadgeId int    `json:"badgeId"`
}

type Arena struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type LeagueStatistics struct {
    CurrentSeason  Season `json:"currentSeason"`
    PreviousSeason Season `json:"previousSeason"`
    BestSeason     Season `json:"bestSeason"`
}

type Season struct {
    ID          string `json:"id,omitempty"`
    Trophies    int    `json:"trophies"`
    BestTrophies int   `json:"bestTrophies"`
}

type Badge struct {
    Name        string `json:"name"`
    Level       int    `json:"level"`
    MaxLevel    int    `json:"maxLevel"`
    Progress    int    `json:"progress"`
    Target      int    `json:"target"`
    IconUrls    IconUrls `json:"iconUrls"`
}

type IconUrls struct {
    Large          string `json:"large,omitempty"`
    Medium         string `json:"medium,omitempty"`
    EvolutionMedium string `json:"evolutionMedium,omitempty"`
}

type Achievement struct {
    Name            string `json:"name"`
    Stars           int    `json:"stars"`
    Value           int    `json:"value"`
    Target          int    `json:"target"`
    Info            string `json:"info"`
    CompletionInfo  *string `json:"completionInfo,omitempty"`
}

type Card struct {
    Name              string `json:"name"`
    ID                int    `json:"id"`
    Level             int    `json:"level"`
    StarLevel         int    `json:"starLevel,omitempty"`
    EvolutionLevel    int    `json:"evolutionLevel"`
    MaxLevel          int    `json:"maxLevel"`
    MaxEvolutionLevel int    `json:"maxEvolutionLevel"`
    Rarity            string `json:"rarity"`
    Count             int    `json:"count"`
    ElixirCost        int    `json:"elixirCost"`
    IconUrls 				IconUrls `json:"iconUrls"`
}

type SeasonResult struct {
		LeagueNumber int     `json:"leagueNumber"`
		Trophies     int     `json:"trophies"`
		Rank         *int    `json:"rank,omitempty"`
}

type RankingPlayerItems struct {
		Tag      string `json:"tag"`
    Name     string `json:"name"`
    ExpLevel int    `json:"expLevel"`
    EloRating int   `json:"eloRating"`
    Rank     int    `json:"rank"`
    Clan     Clan   `json:"clan"`
}

type Cursors struct {
	Before *string `json:"before,omitempty"`
	After  *string `json:"after,omitempty"`
}
type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type RankingPlayersResponse struct {
  Items []RankingPlayerItems `json:"items"`
	Paging Paging `json:"paging"`
}

type PlayerData struct {
	Wins                            int                     `json:"wins"`
	Losses                          int                     `json:"losses"`
	CurrentPathOfLegendSeasonResult PathOfLegendSeasonResult `json:"currentPathOfLegendSeasonResult"`
	LastPathOfLegendSeasonResult    PathOfLegendSeasonResult `json:"lastPathOfLegendSeasonResult"`
	BestPathOfLegendSeasonResult    PathOfLegendSeasonResult `json:"bestPathOfLegendSeasonResult"`
}

type PathOfLegendSeasonResult struct {
	Trophies     int `json:"trophies"`
	LeagueNumber int `json:"leagueNumber"`
	Ranking      int `json:"ranking,omitempty"`
}

type DailyRankingLog struct {
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	Rank      int    `json:"rank"`
	EloRating int    `json:"eloRating"`
}