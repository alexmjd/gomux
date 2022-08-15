package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating the Basic structure
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// Creating a new type "allEvents" of type "Slice of Event" (kind of an alias)
type allEvents []event

// Creating a new variable of type "allEvents" which contain a Json of event
var events = allEvents{
	{
		ID:          "1",
		Title:       "API Tutorial",
		Description: "Making a tutorial to learn how to build a Golang API",
	},
}

// Create a new event (A Request handler take always 2 args, the ResponseWritter and a pointer on Request)
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "A problem occured")
	}

	// Put the content of "reqBody" in our "newEvent" variable, passing its reference
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)

	// Set the ResponseWriter's Header to Status Created (201)
	w.WriteHeader(http.StatusCreated)

	// Write in the ResponseWriter the encoded newEvent
	json.NewEncoder(w).Encode(newEvent)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent)

	log.Fatal(http.ListenAndServe(":8088", router))
}
