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
	dbHandler, _ := dblayer.NewPersistenceLayer(config.DatabaseType, config.DBConnection)

	fmt.Println("type: %v, connection: %v", config.DatabaseType, config.DBConnection)

	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbHandler))
}