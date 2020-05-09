package team

//Player datatype for Collegiate CS players
type Player struct {
	SteamID  int    `json:"SteamID"`  //Every player has a STEAM ID given by valve. Can be used for information using Valve's API.
	TeamID   int    `json:"TeamID"`   //Used to link a player to a team. And for algorithm logic.
	Username string `json:"Username"` //Username is the players in-game name.
	Uni      string `json:"Uni"`      //University name they are associated with.
	UniID    string `json:"UniID"`    //University ID they are associated with.
	Stats    Stats  `json:"Stats"`    //Stats block primarily used for webservices.
	//May require more information aka Academic information and a data type of associated media of the player.
}

//Players is a list of players. University roster list
type Players []Player

//Roster is a slice of player aka a team roster
type Roster struct {
	TeamID     int      `json:"TeamID"`     //Used to link a player to a team. And for algorithm logic.
	TeamName   string   `json:"TeamName"`   //Players enjoy naming their teams :)
	Uni        string   `json:"Uni"`        //University name they are associated with.
	UniID      string   `json:"UniID"`      //University ID they are associated with.
	TeamRoster []Player `json:"TeamRoster"` //List of player objects of the roster.
	//teamName and Uni is redundant for webservice integration. AKA Player stat webpage will use player.Uni() && while team match webpage will use team.Uni()
}

//Rosters is a set of Roster. Used for University Roster pages
type Rosters []Roster

//Stats is a datastructure for the stats so that the player structure is not filled with a ton of non-player specific values
type Stats struct {
	GamesPlayed int `json:"GamesPlayed"` //Games played
	GameWins    int `json:"GameWins"`    //Game wins
	GameLoss    int `json:"GameLoss"`    //Game Loss
	GameTie     int `json:"GameTie"`     //Game Tie (if applicable)
	//Lots more, but this is fine for testing
}

//allStats was added in case I need it in the future
type allStats []Stats
