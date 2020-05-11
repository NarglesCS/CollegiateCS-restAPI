package main

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

//Player datatype for Collegiate CS players
type Player struct {
	SteamID  string `json:"SteamID"`  //Every player has a STEAM ID given by valve. Can be used for information using Valve's API.
	TeamID   int    `json:"TeamID"`   //Used to link a player to a team. And for algorithm logic.
	Username string `json:"Username"` //Username is the players in-game name.
	Uni      string `json:"Uni"`      //University name they are associated with.
	UniID    int    `json:"UniID"`    //University ID they are associated with.
	Captain  bool   `json:"Captain"`  //Team captain True/False
	Stats    Stats  `json:"Stats"`    //Stats block primarily used for webservices.
	//May require more information aka Academic information and a data type of associated media of the player.
}

//Players is a list of players. University roster list
type Players []Player

//buildPlayers will build a list of Players from type *[]string of teamID
func buildPlayers(s *[]string) (*Rosters, *[]int) {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()
	var lsdbPlayers []dbPlayer
	var lsPlayers Players
	var lsRoster Rosters
	var lstID []int
	for _, id := range *s {

		if intID, err := strconv.Atoi(id); err == nil {
			db.Where("team_id = ?", intID).Find(&lsdbPlayers)

			for _, pro := range lsdbPlayers {
				pStats := connectStats(&pro)
				rPlayer := Player{
					SteamID:  pro.SteamID,  //Every player has a STEAM ID given by valve. Can be used for information using Valve's API.
					TeamID:   pro.TeamID,   //Used to link a player to a team. And for algorithm logic.
					Username: pro.Username, //Username is the players in-game name.
					Uni:      pro.Uni,      //University name they are associated with.
					UniID:    pro.UniID,    //University ID they are associated with.
					Stats:    *pStats,      //Stats block primarily used for webservices.
				}
				lsPlayers = append(lsPlayers, rPlayer)
			}
			pRos := buildRoster(&lsPlayers, intID)
			lsRoster = append(lsRoster, *pRos)
			lstID = append(lstID, intID)
		}
	}
	return &lsRoster, &lstID
}

//DummyPlayer returns an example Player object
func DummyPlayer() Player {
	ExPlayer := Player{
		SteamID:  "STEAM_0:1:32729848",        //Every player has a STEAM ID given by valve. Can be used for information using Valve's API.
		TeamID:   123,                         //Used to link a player to a team. And for algorithm logic.
		Username: "Nargles",                   //Username is the players in-game name.
		Uni:      "Michigan State University", //University name they are associated with.
		UniID:    1,                           //University ID they are associated with.
		Stats:    DummyStats(),                //Stats block primarily used for webservices.
	}
	return ExPlayer
}
