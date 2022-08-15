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

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	// Get the id argument from the request
	eventId := mux.Vars(r)["id"]

	// Look in events slice if there's a match
	for _, singleEvent := range events {
		if singleEvent.ID == eventId {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

	// Get the id from request
	eventId := mux.Vars(r)["id"]
	var updatedEvent event

	// Read all data in body and put them in reqBody
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Cannot update, error occured.\n")
	}

	// Put reqBody json contents inside the event updatedEvent (given by reference)
	json.Unmarshal(reqBody, &updatedEvent)

	// Go throught slice of events
	for i, singleEvent := range events {
		if singleEvent.ID == eventId {

			// Change values if found
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description

			// Change the value of event at index in the slice by the new updatedEvent
			events = append(events[:i], singleEvent)

			// Write the new updatedEvent to the ResponseWriter
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8088", router))
}
