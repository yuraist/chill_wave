package main

import (
	"chill_wave/eventservice/rest"
	"chill_wave/lib/configuration"
	"chill_wave/lib/persistence/dblayer"
	"flag"
	"fmt"
	"log"
)

func main() {
	confPath := flag.String("conf", `../lib/configuration/config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbHandler, err := dblayer.NewPersistenceLayer(config.DatabaseType, config.DBConnection)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful...")

	httpErrorChan, httpTLSErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndpoint, dbHandler)

	// Block the main goroutine to wait for multiple channels
	select {
	case err := <- httpErrorChan:
		log.Fatal("HTTP Error: ", err)
	case err := <- httpTLSErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}