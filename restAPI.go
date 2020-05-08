package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Game defines a new data structure for a game
type Game struct {
	ID       int      `json:"id"`
	Teams    []string `json:"Teams"`
	Players1 []string `json:"Players1"`
	Players2 []string `json:"Players2"`
}

//Games initialises a list of games
type Games []Game

func allGames(w http.ResponseWriter, r *http.Request) {
	//Construct an article object
	games := Games{
		Game{
			ID:       1,
			Teams:    []string{"MSU", "UM"},
			Players1: []string{"Nargles", "Congala", "Keeb", "Zeeker", "Sensed"},
			Players2: []string{"Angel", "Prophet", "Psych", "Koi", "Fidel"},
		},
	}
	//Print article object in the form of JSON
	fmt.Println("Endpoint hit: All articles endpoint")
	json.NewEncoder(w).Encode(games)
}

//POST request handle for inputting future games. Will link to SQL for full integration with GET information pulls.
func addGames(w http.ResponseWriter, r *http.Request) {
	//Pull the useful information **add even handling for content type != "application/json"**
	req := r.Body

	// convert body to a byte stream
	body, err := ioutil.ReadAll(req)

	//if no errors echo response back to sender in their console (debugging stuff i guess)
	if err == nil && body != nil {
		fmt.Fprintf(w, string(body))
	}
	//Create a new game instance for consistency
	var nGame Game

	//Unmarshal the body and write to the referenced variable at the same time.
	errr := json.Unmarshal(body, &nGame)
	if errr != nil {
		//Insert error handling statements here.
	}
	// Print statement to see if i can access the info. (I can :)
	fmt.Fprintf(w, string(nGame.Teams[0]))
}

//Function to display a homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

//Request handler function for the
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Schedule", allGames).Methods("GET")
	myRouter.HandleFunc("/Schedule", addGames).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
