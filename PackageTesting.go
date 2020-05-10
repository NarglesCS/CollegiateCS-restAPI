package main

import (
	"game"
)

func main() {
	exGame := game.DummyGame() //GameID, Teams[2], TeamsID[2], Map,Roster1, Roster2
	exSet := game.DummySet()   //MatchID, Teams[2], TeamsID[2], MapList[3], GameList[3]
	//exStats := team.DummyStats() 		//GamesPlayed, GameWins, GameLoss, GameTie
	//exPlayer := team.DummyPlayer() 	//SteamID, TeamID, Username, Uni, UniID, Stats
	//exRoster := team.DummyRoster() 	//TeamID, TeamName, Uni, UniID, TeamRoster

	//wat := exGame.Roster(exGame.TeamsID[1])
	game.Print(&exGame)
	game.SetPrint(&exSet)
}
