package game

import (
	"fmt"
	"strconv"
	"strings"
	"team"
)

var person team.Player

//Game defines a new data structure for a single game (Bo1)
type Game struct {
	GameID  int         `json:"GameID"`
	Teams   [2]string   `json:"Teams"`
	TeamsID [2]int      `json:"TeamsID"`
	Map     string      `json:"Map"`
	Roster1 team.Roster `json:"Roster1"`
	Roster2 team.Roster `json:"Roster2"`
}

//Games is a list of games. Added just in case but the only use should be in Setmatch stuct
type Games []Game

//Set initialises a Bo3 or Bo5
type Set struct {
	MatchID  int       `json:"MatchID"`
	Teams    [2]string `json:"Teams"`
	TeamsID  [2]int    `json:"TeamsID"`
	MapList  []string  `json:"MapList"`
	GameList []Game    `json:"GameList"`
}

//Sets was created just in case i would need it. But will probably be unused ~~Unused atm~~
type Sets []Set

//GameNum method to get id value
func (g *Game) GameNum() int {
	return (g.GameID)
}

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

//Roster method to report players based on team selection (TEAM 1 should always equal Players 1) ~~~Testing ATM~~~
func (g *Game) Roster(tmID int) team.Roster {
	var roster team.Roster
	if tmID == g.Roster1.TeamID {
		roster = g.Roster1
	} else if tmID == g.Roster2.TeamID {
		roster = g.Roster2
	}
	// returns roster 1 by default until I build a teamID == teamID logic statment
	return roster
}

//Match method to report teams for Setmatch datatype
func (s *Set) Match() []string {
	var matches []string
	for _, item := range s.GameList {
		matches = append(matches, item.Match())
	}
	return (matches)
}

//GameNum for type Set
func (s *Set) GameNum() []int {
	var id []int
	for _, item := range s.GameList {
		id = append(id, item.GameNum())
	}
	return (id)
}

//Output a []string type as a single string deliminated with a " " character.
func Output(ls []string) string {
	return strings.Join(ls, ", ")
}

//Print is a function to output a game object.
func Print(g *Game) {
	fmt.Println("Game ID: " + strconv.Itoa(g.GameID))
	fmt.Println("Teams: " + g.Match())
	fmt.Println(g.Team1() + " Roster: " + Output(g.Roster1.PlayerList()))
	fmt.Println(g.Team2() + " Roster: " + Output(g.Roster1.PlayerList()))
	fmt.Println()
}

//SetPrint is a function to output a Setmatch object
func SetPrint(s *Set) {
	for _, g := range s.GameList {
		Print(&g)
	}
}

//DummyGame returns an example Game object
func DummyGame() Game {
	exGame := Game{
		GameID:  1,
		Teams:   [2]string{"MSU", "UM"},
		TeamsID: [2]int{123, 890},
		Map:     "DE_NUKE",
		Roster1: team.DummyRoster(),
		Roster2: team.DummyRoster(),
	}
	return exGame
}

//DummySet returns an example Setmatch object
func DummySet() Setmatch {
	exSet := Setmatch{
		MatchID: 5423,
		Teams:   [2]string{"MSU", "UM"},
		TeamsID: [2]int{123, 890},
		MapList: []string{"DE_NUKE", "DE_OVERPASS", "DE_MIRAGE"},
		GameList: []Game{
			DummyGame(),
			DummyGame(),
			DummyGame(),
		},
	}
	return exSet
}
