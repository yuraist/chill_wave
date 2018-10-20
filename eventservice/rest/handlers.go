package rest

import (
	"fmt"
	"net/http"
)

type eventServiceHandler struct {}

func (eventHandler *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "find event")
}

func (eventHandler *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all events")
}

func (eventHandler *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a new event")
}
