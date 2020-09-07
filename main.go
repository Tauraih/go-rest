package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome home!")
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Decsription"`
}

type allEvents []event

var events = allEvents{
	{
		ID:             "1",
		Title:          "Intro to Golang",
		Description:    "Come join us",
	},
	{
		ID:             "2",
		Title:          "Intro to Golang",
		Description:    "Come join us",
	},
	{
		ID:             "3",
		Title:          "Intro to Golang",
		Description:    "Come join us",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqbody, err := ioutil.ReadAll(r.Body)
	if err!= nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and descrption only")
	}

	json.Unmarshal(reqbody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updateEvent event

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data with event title & description")
	}
	json.Unmarshal(reqbody, &updateEvent)

	for i, singleEvent := range events{
		if singleEvent.ID == eventID {
			singleEvent.Title = updateEvent.Title
			singleEvent.Description = updateEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", index)
	api.HandleFunc("/event", createEvent).Methods("POST")
	api.HandleFunc("/events", getAllEvents).Methods("GET")
	api.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	api.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	api.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", router))
}
