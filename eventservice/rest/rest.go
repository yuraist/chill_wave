package rest

import (
	"chill_wave/lib/persistence"
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := newEventHandler(dbHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events")

	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	return http.ListenAndServe(endpoint, r)
}