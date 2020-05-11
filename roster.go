package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Roster is a slice of player aka a team roster
type Roster struct {
	TeamID     int     `json:"TeamID"`     //Used to link a player to a team. And for algorithm logic.
	TeamName   string  `json:"TeamName"`   //Players enjoy naming their teams :)
	Uni        string  `json:"Uni"`        //University name they are associated with.
	UniID      int     `json:"UniID"`      //University ID they are associated with.
	TeamRoster Players `json:"TeamRoster"` //List of player objects of the roster.
	//teamName and Uni is redundant for webservice integration. AKA Player stat webpage will use player.Uni() && while team match webpage will use team.Uni()
}

//Rosters is a set of Roster. Used for University Roster pages
type Rosters []Roster

//PlayerList returns a list of all players on a given roster
func (r *Roster) PlayerList() []string {
	var lsPlayer []string
	for _, pro := range r.TeamRoster {
		lsPlayer = append(lsPlayer, pro.Username)
	}

	return lsPlayer
}

//buildRoster builds the rosters associated with the match.
func buildRoster(p *Players, tID int) *Roster {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()
	defer db.Close()
	var lsdbRoster dbRoster
	var pRoster Roster
	db.Where("team_id = ?", tID).Find(&lsdbRoster)

	pRoster = Roster{
		TeamID:     lsdbRoster.TeamID,   //Used to link a player to a team. And for algorithm logic.
		TeamName:   lsdbRoster.TeamName, //Players enjoy naming their teams :)
		Uni:        lsdbRoster.Uni,      //University name they are associated with.
		UniID:      lsdbRoster.UniID,    //University ID they are associated with.
		TeamRoster: *p,
	}
	return &pRoster
}

//DummyRoster returns an example Roster datatype
func DummyRoster() Roster {
	ExRoster := Roster{
		TeamID:   123,                         //Used to link a player to a team. And for algorithm logic.
		TeamName: "MSU Spartans",              //Players enjoy naming their teams :)
		Uni:      "Michigan State University", //University name they are associated with.
		UniID:    1,                           //University ID they are associated with.
		TeamRoster: Players{
			DummyPlayer(),
			DummyPlayer(),
			DummyPlayer(),
			DummyPlayer(),
			DummyPlayer(),
		},
	}
	return ExRoster
}
