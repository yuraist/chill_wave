package main

import (
	"chill_wave/eventservice/rest"
	"chill_wave/lib/configuration"
	msgqueue_amqp "chill_wave/lib/msgqueue/amqp"
	"chill_wave/lib/persistence/dblayer"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	confPath := flag.String("conf", `../lib/configuration/config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)
	conn, err := amqp.Dial(config.AMQPMessageBroker)

	if err != nil {
		panic(err)
	}

	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)

	if err != nil {
		panic(err)
	}

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