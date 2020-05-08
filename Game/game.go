package game

//Game defines a new data structure for a game
type Game struct {
	ID       int      `json:"id"`
	Teams    []string `json:"Teams"`
	Players1 []string `json:"Players1"`
	Players2 []string `json:"Players2"`
}

//Games initialises a list of games
type Games []Game
