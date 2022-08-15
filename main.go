package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "API Tutorial",
		Description: "Making a tutorial to learn how to build a Golang API",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	log.Fatal(http.ListenAndServe(":8088", router))
}
