package rest

import (
	"chill_wave/lib/persistence"
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI(endpoint string, tlsendpoint string, dbHandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := newEventHandler(dbHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()

	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	httpErrorChannel := make(chan error)
	httpTLSErrorChannel := make(chan error)

	go func() {
		httpErrorChannel <- http.ListenAndServe(endpoint, r)
	}()

	go func() {
		httpTLSErrorChannel <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r)
	}()

	return httpErrorChannel, httpTLSErrorChannel
}