package main

import (
	"fmt"
	"game"
)

//iterate through []Game datatypes
func iterGames(g game.Games) {
	for _, item := range g {
		fmt.Println(item.Match())
	}
}

func main() {
	games := game.Games{
		game.Game{
			ID:       1,
			Teams:    []string{"MSU", "UM"},
			Players1: []string{"Nargles", "Congala", "Keeb", "Zeeker", "Sensed"},
			Players2: []string{"Angel", "Prophet", "Psych", "Koi", "Fidel"},
		},
		game.Game{
			ID:       2,
			Teams:    []string{"WMU", "DU"},
			Players1: []string{"Nargles", "Congala", "Keeb", "Zeeker", "Sensed"},
			Players2: []string{"Angel", "Prophet", "Psych", "Koi", "Fidel"},
		},
	}
	singleGame := game.Game{
		ID:       1,
		Teams:    []string{"MSU", "UM"},
		Players1: []string{"Nargles", "Congala", "Keeb", "Zeeker", "Sensed"},
		Players2: []string{"Angel", "Prophet", "Psych", "Koi", "Fidel"},
	}
	iterGames(games)
	fmt.Println("Game Datatype")
	fmt.Println(singleGame.GameNum())
	fmt.Println(singleGame.Match())
	fmt.Println(singleGame.Team1())
	fmt.Println(singleGame.Team2())
	fmt.Println(singleGame.Players(1))
	fmt.Println(singleGame.Players(2))
	fmt.Println(singleGame.Players(3))
	fmt.Println("Games Datatype")
	fmt.Println(games.GameNum())
	fmt.Println(games.Match())
	//fmt.Println(games.Team1())
	//fmt.Println(games.Team2())
	//fmt.Println(games.Players(1))
	//fmt.Println(games.Players(2))
	//fmt.Println(games.Players(3))

}
