package main

//Set initialises a Bo3 or Bo5
type Set struct {
	MatchID   int       `json:"MatchID"`
	MatchDate string    `json:"MatchDate"`
	Teams     [2]string `json:"Teams"`
	TeamsID   [2]int    `json:"TeamsID"`
	MapList   []string  `json:"MapList"`
	GameList  []Game    `json:"GameList"`
}

//Sets was created just in case i would need it. But will probably be unused ~~Unused atm~~
type Sets []Set

/*
//Match method to report teams for Setmatch datatype
func (s *Set) Match() []string {
	var matches []string
	for _, item := range s.GameList {
		matches = append(matches, item.Match())
	}
	return (matches)
}
*/

//GameNum for type Set
func (s *Set) GameNum() []int {
	var id []int
	for _, item := range s.GameList {
		id = append(id, item.GameNum())
	}
	return (id)
}

/*
//SetPrint is a function to output a Setmatch object
func SetPrint(s *Set) {
	for _, g := range s.GameList {
		Print(&g)
	}
}
*/

//DummySet returns an example Setmatch object
func DummySet() Set {
	exSet := Set{
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
