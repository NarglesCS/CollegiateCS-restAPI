package main

import (
	"encoding/json"
	"fmt"
	"game"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Function to return all games in a db(not hooked up or coded)
func allGames(w http.ResponseWriter, r *http.Request) {
	/*Code to access SQL
	 */

	/* Code for iterating through rows
	gameRows := db.Query("select x,y,z from whatever")
	var lsGames game.Games
	var curGame game.Game
	for rows.Next(){

		if err := gameRows.Scan(&curGame.GameID, &curGame.Teams, &curGame.TeamsID, &curGame.Map, &curGame.Roster1, &curGame.Roster2){
			Error code here (NOTE: Roster1 and Roster2 will be more complicated because they are a custom object)
		} else {
			lsGames = append(lsGames, curGame)
		}
	}
	*/

	//Construct an game object
	MSUset := game.DummySet()
	//Print game object in the form of JSON
	fmt.Println("Endpoint hit: All games endpoint")
	json.NewEncoder(w).Encode(MSUset)
}

//POST request handle for inputting future games. Will link to SQL for full integration with GET information pulls.
func addGames(w http.ResponseWriter, r *http.Request) {
	//Check to see if the content type is == "application/json"
	reqHead := r.Header
	if reqHead.Get("Content-Type") == "application/json" {
		println("JSON baby")

		//Pull the useful information **add even handling for content type != "application/json"**
		req := r.Body

		// convert body to a byte stream
		body, err := ioutil.ReadAll(req)

		//if no errors echo response back to sender in their console (debugging stuff i guess)
		if err == nil && body != nil {
			fmt.Fprintf(w, string(body))
		}
		//Create a new game instance for consistency
		var nGame game.Game

		//Unmarshal the body and write to the referenced variable at the same time.
		errr := json.Unmarshal(body, &nGame)
		if errr != nil {
			//Insert error handling statements here.
		}
		/*Code for SQL connection and INSERT query.
		 */

	} else {
		fmt.Fprintf(w, "Post Request not of type application/json")
	}

}

//showMatch is designed to return a specific game for the showmatch part of webpage
func showMatch(w http.ResponseWriter, r *http.Request) {
	/*Code to access SQL
	 */

	/* Code for iterating through rows
	gameRows := db.Query("select x,y,z from whatever")
	for rows.Next(){
		var showMatch game.Game
		if err := gameRows.Scan(&showMatch.GameID, &showMatch.Teams, &showMatch.TeamsID, &showMatch.Map, &showMatch.Roster1, &showMatch.Roster2){
			Error code here (NOTE: Roster1 and Roster2 will be more complicated because they are a custom object)
		}
		else if showMatch.GameID == 1214{
			json.NewEncoder(w).Encode(showMatch)
		}
	}
	*/
}
func gameUpcoming(w http.ResponseWriter, r *http.Request) {
	/*Code to access SQL
	 */

	/* Code for iterating through rows
	gameRows := db.Query("select x,y,z from whatever")
	var lsGames game.Games
	var curGame game.Game
	for rows.Next(){

		if err := gameRows.Scan(&curGame.GameID, &curGame.Teams, &curGame.TeamsID, &curGame.Map, &curGame.Roster1, &curGame.Roster2){
			Error code here (NOTE: Roster1 and Roster2 will be more complicated because they are a custom object)
		} else {
			lsGames = append(lsGames, curGame)
		}
	}
	*/
	/*Code to find the latest 5 additions to the game list.
	var upcomingGames game.Games
	for i=0;i<5{
		upcomingGames = append(upcomingGames, lsGames[i])
		i++
	}
	*/
}

func showResults(w http.ResponseWriter, r *http.Request) {
	/*Code to access SQL
	 */

	/* Code for iterating through rows
	resultRows := db.Query("select x,y,z from whatever")
	var lsResults game.Results
	var curResult game.Result
	for rows.Next(){

		if err := resultRows.Scan(&curResult.GameID, &curResult.Teams, &curResult.TeamsID, &curResult.Map, &curResult.Roster1, &curResult.Roster2){
			Error code here (NOTE: Roster1 and Roster2 will be more complicated because they are a custom object)
		} else {
			lsResults = append(lsResults, curcurResultGame)
		}
	}
	json.NewEncoder(w).Encode(lsResults)
	*/

}

//Function to display a homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

//Request handler function for the
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Schedule", allGames).Methods("GET")     //Service for upcoming games page 				-Ready for SQL testing
	myRouter.HandleFunc("/Upcoming", gameUpcoming).Methods("GET") //Service for the next 5 upcoming games			-Ready for SQL testing
	myRouter.HandleFunc("/Showmatch", showMatch).Methods("GET")   //Service for the showmatch						-Ready for SQL testing
	myRouter.HandleFunc("/Schedule", addGames).Methods("POST")    //Service to put upcoming games in the system		-Ready for SQL testing
	myRouter.HandleFunc("/Results", showResults).Methods("GET")   //Service for the results page					-Not ready

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
