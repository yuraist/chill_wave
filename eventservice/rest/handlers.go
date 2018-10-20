package rest

import (
	"chill_wave/lib/persistence"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func newEventHandler(databaseHandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{dbhandler: databaseHandler,}
}

func (eventHandler *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]

	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either
										search by id via /id/4 
										to search by name via /name/coldplayconcert}`)
		return
	}

	searchKey, ok := vars["search"]

	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either
										search by id via /id/4 
										to search by name via /name/coldplayconcert}`)
		return
	}

	var event persistence.Event
	var err error

	switch strings.ToLower(criteria) {
	case "name":
		event, err = eventHandler.dbhandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eventHandler.dbhandler.FindEvent(id)
		}
	}

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "{error %s}", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eventHandler *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eventHandler.dbhandler.FindAllAvailableEvents()
	fmt.Println("All events start")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to find all available events %s}", err)
		return
	}

	fmt.Println("All events mid")
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while trying to encode events to JSON %s}", err)
	}

	fmt.Println("All events end")
}

func (eventHandler *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	var event persistence.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while decoding event data %s}", err)
		return
	}

	id, err := eventHandler.dbhandler.AddEvent(event)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occured while persisting event %d %s}", id, err)
		return
	}

	fmt.Fprintf(w, `{"id": %d}`, id)
}
