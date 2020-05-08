package game

import "strings"

//Game defines a new data structure for a game
type Game struct {
	ID       int      `json:"id"`
	Teams    []string `json:"Teams"`
	Players1 []string `json:"Players1"`
	Players2 []string `json:"Players2"`
}

//Games initialises a list of games
type Games []Game

//Match method to report teams ({TEAM 1} vs {TEAM 2})
func (g *Game) Match() string {
	return (string(g.Teams[0]) + " vs. " + string(g.Teams[1]))
}

//Team1 method to report team1
func (g *Game) Team1() string {
	return (string(g.Teams[0]))
}

//Team2 method to report team2
func (g *Game) Team2() string {
	return (string(g.Teams[1]))
}

//Players method to report players based on team selection (TEAM 1 should always equal Players 1)
func (g *Game) Players(team int) []string {
	if team == 1 {
		return (g.Players1)
	} else if team == 2 {
		return (g.Players2)
	}
	//The ... is important because append is a variadic function. Research says thats any function that accepts variable amounts of inputs.
	//It seems that since append normally accepts a slice and a single input or several single inputs. There is no issue.
	//This is because it collects the individual inputs and creates a new slice.
	//But, the ... denotes that it is already a slice and passes it on to the other methods behind the scenes. Or else there will be an error.
	return (append(g.Players1, g.Players2...))
}

//Output a []string type as a single string deliminated with a " " character.
func (g *Game) Output(ls []string) string {

	return strings.Join(ls, " ")
}
