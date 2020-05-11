package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	//justify this
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//DecodeSQL decodes the SQL outputs by splitting at the "-" because SQL doesn't allow certain things in the db how it is reassembled is a case by case thing.
//Which will be done in the function operating on the value
func DecodeSQL(s *string) *[]string {
	ls := strings.Split(*s, "-")
	return &ls
}

func allStats(c chan []dbStats) {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()

	var dbStats []dbStats
	db.Find(&dbStats)
	c <- dbStats
}
func allPlayers(c chan []dbPlayer) {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()

	var dbPlayers []dbPlayer
	db.Find(&dbPlayers)
	c <- dbPlayers
}
func allRosters(c chan []dbRoster) {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()

	var dbRosters []dbRoster
	db.Find(&dbRosters)
	c <- dbRosters
}

//AllGames is a sql connected allGames
func AllGames(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	//cStats := make(chan []dbStats)
	//cPlayers := make(chan []dbPlayer)
	//cRosters := make(chan []dbRoster)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()
	//Trigger happy to create go routines, alwasy wanted to use them
	//go allStats(cStats)
	//go allPlayers(cPlayers)
	//go allRosters(cRosters)

	var dbGames []dbGame
	db.Find(&dbGames)
	//dbStatsOp := <-cStats
	//dbPlayers := <-cPlayers
	//dbRosters := <-cRosters
	//close(cStats)
	//close(cPlayers)
	//close(cRosters)
	aGames := buildGame(&dbGames)
	json.NewEncoder(w).Encode(*aGames)
}

//dbPlayer is db Friendly player
type dbPlayer struct {
	SteamID  string `json:"SteamID"`  //Every player has a STEAM ID given by valve. Can be used for information using Valve's API.
	TeamID   int    `json:"TeamID"`   //Used to link a player to a team. And for algorithm logic.
	Username string `json:"Username"` //Username is the players in-game name.
	Uni      string `json:"Uni"`      //University name they are associated with.
	UniID    int    `json:"UniID"`    //University ID they are associated with.
	Captain  bool   `json:"Captain"`  //Team captain True/False

	//Stats    Stats  `json:"Stats"`    //Stats block primarily used for webservices. Manually added later.
	//May require more information aka Academic information and a data type of associated media of the player.
}

//dbRoster is db friendly roster
type dbRoster struct {
	TeamID   int    `json:"TeamID"`   //Used to link a player to a team. And for algorithm logic.
	TeamName string `json:"TeamName"` //Players enjoy naming their teams :)
	Uni      string `json:"Uni"`      //University name they are associated with.
	UniID    int    `json:"UniID"`    //University ID they are associated with.
	//TeamRoster []Player `json:"TeamRoster"` //List of player objects of the roster.
	//teamName and Uni is redundant for webservice integration. AKA Player stat webpage will use player.Uni() && while team match webpage will use team.Uni()
}

//dbGame is a db friendly game
type dbGame struct {
	GameID    int    `json:"GameID"`
	MatchDate string `json:"MatchDate"`
	MatchID   int    `json:"MatchID"`
	//Teams     string `json:"Teams"`   //array of string ~~Convert from [2]string to string delimination ~~ommitted in favor of a different approach
	TeamsID string `json:"TeamsID"` //array of string ~~Convert from [2]int to string delimination	~~Not Done
	Map     string `json:"Map"`
	//	Roster1 team.Roster `json:"Roster1"`
	//	Roster2 team.Roster `json:"Roster2"`
}

//dbSet is a db friendly Set
type dbSet struct {
	MatchID   int    `json:"MatchID"`
	MatchDate string `json:"MatchDate"`
	//Teams     string `json:"Teams"`   //array of string ~~Convert from [2]string to string delimination	~~ommitted in favor of a different approach
	TeamsID string `json:"TeamsID"` //array of string ~~Convert from [2]int to string delimination	~~Not Done
	MapList string `json:"MapList"` //array of string ~~Convert from []string to string delimination	~~Not Done
	//GameList []Game    `json:"GameList"`
}

//dbStat struct created for when Stats become db unfriendly
type dbStats struct {
	SteamID     string `json:"SteamID"`     //Steam ID to tie to player
	GamesPlayed int    `json:"GamesPlayed"` //Games played
	GameWins    int    `json:"GameWins"`    //Game wins
	GameLoss    int    `json:"GameLoss"`    //Game Loss
	GameTie     int    `json:"GameTie"`     //Game Tie (if applicable)
	//Lots more, but this is fine for testing
}

//dbUni struct for uniID finding

//InitialMigration does something
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "collegiateCS.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to db")
	}
	defer db.Close()

	//db.AutoMigrate(&dbStats{})
	//db.AutoMigrate(&dbPlayer{})
	//db.AutoMigrate(&dbRoster{})
	//db.AutoMigrate(&dbGame{})
	//db.AutoMigrate(&dbSet{})
}
