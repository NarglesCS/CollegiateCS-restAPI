package main

import (
	"strings"
)

//Game defines a new data structure for a single game (Bo1)
type Game struct {
	GameID    int    `json:"GameID"`
	MatchDate string `json:"MatchDate"`
	MatchID   int    `json:"MatchID"`
	//Teams     []string `json:"Teams"`
	TeamsID []int   `json:"TeamsID"`
	Map     string  `json:"Map"`
	Rosters Rosters `json:"Rosters"`
}

//Games is a list of games. Added just in case but the only use should be in Setmatch stuct
type Games []Game

//DecodeSQL will decode the team ID and return the roster information

//GameNum method to get id value
func (g *Game) GameNum() int {
	return (g.GameID)
}

/*
//Match method to report teams ({TEAM 1} vs {TEAM 2})
func (g *Game) Match() string {
	return (string(g.Teams[0]) + " vs. " + string(g.Teams[1]) + " on " + g.Map)
}

//Team1 method to report team1
func (g *Game) Team1() string {
	return (string(g.Teams[0]))
}

//Team2 method to report team2
func (g *Game) Team2() string {
	return (string(g.Teams[1]))
}
*/

//Roster method to report players based on team selection (TEAM 1 should always equal Players 1) ~~~Testing ATM~~~
func (g *Game) Roster(tmID int) Roster {
	var roster Roster
	if tmID == g.Rosters[0].TeamID {
		roster = g.Rosters[0]
	} else if tmID == g.Rosters[1].TeamID {
		roster = g.Rosters[1]
	}
	// returns roster 1 by default until I build a teamID == teamID logic statment
	return roster
}

//Output a []string type as a single string deliminated with a " " character.
func Output(ls []string) string {
	return strings.Join(ls, ", ")
}

/*
//Print is a function to output a game object.
func Print(g *Game) {
	fmt.Println("Game ID: " + strconv.Itoa(g.GameID) + " Match ID: " + strconv.Itoa(g.MatchID))
	fmt.Println("Teams: " + g.Match())
	fmt.Println(g.Team1() + " Roster: " + Output(g.Rosters[0].PlayerList()))
	fmt.Println(g.Team2() + " Roster: " + Output(g.Rosters[1].PlayerList()))
	fmt.Println()
}
*/

//DummyGame returns an example Game object
func DummyGame() Game {
	exGame := Game{
		GameID:  1,
		MatchID: 12,
		//Teams:   [2]string{"MSU", "UM"},
		TeamsID: []int{123, 890},
		Map:     "DE_NUKE",
		Rosters: Rosters{
			DummyRoster(),
			DummyRoster(),
		},
	}
	return exGame
}

//buildGame assembles a Game object
func buildGame(g *[]dbGame) *Games {
	var lsGames Games

	for _, game := range *g {
		plsGameID := DecodeSQL(&game.TeamsID)
		pRoster, plstID := buildPlayers(plsGameID)
		gGame := Game{
			GameID:  game.GameID,
			MatchID: game.MatchID,
			TeamsID: *plstID,
			Map:     game.Map,
			Rosters: *pRoster,
		}
		lsGames = append(lsGames, gGame)
	}

	return &lsGames
}
