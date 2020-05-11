package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Stats is a datastructure for the stats so that the player structure is not filled with a ton of non-player specific values
type Stats struct {
	SteamID     string `json:"SteamID"`     //Steam ID to tie to player
	GamesPlayed int    `json:"GamesPlayed"` //Games played
	GameWins    int    `json:"GameWins"`    //Game wins
	GameLoss    int    `json:"GameLoss"`    //Game Loss
	GameTie     int    `json:"GameTie"`     //Game Tie (if applicable)
	//Lots more, but this is fine for testing
}

//allStats was added in case I need it in the future
type allPlayerStats []Stats

//DummyStats is supposed to return an example Stats object
func DummyStats() Stats {
	ExStats := Stats{
		SteamID:     "1",
		GamesPlayed: 100, //Games played
		GameWins:    69,  //Game wins
		GameLoss:    30,  //Game Loss
		GameTie:     1,   //Game Tie (if applicable)
	}
	return ExStats
}

//connectStats finds the stats and returns them
func connectStats(pro *dbPlayer) *Stats {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()
	var lsdbStats dbStats
	var pStats Stats
	db.Where("steam_id = ?", pro.SteamID).Find(&lsdbStats)
	pStats = Stats{
		SteamID:     lsdbStats.SteamID,     //Steam ID to tie to player
		GamesPlayed: lsdbStats.GamesPlayed, //Games played
		GameWins:    lsdbStats.GameWins,    //Game wins
		GameLoss:    lsdbStats.GameLoss,    //Game Loss
		GameTie:     lsdbStats.GameTie,     //Game Tie (if applicable)
	}

	return &pStats
}
