package configuration

import (
	"chill_wave/lib/persistence/dblayer"
	"encoding/json"
	"fmt"
	"os"
)

var (
	DBTypeDefault          = dblayer.DBTYPE("mongodb")
	DBConnectionDefault    = "mongodb://127.0.0.1"
	RestfulEndpointDefault = "localhost:8181"
	RestfulTLSEndpointDefault = "localhost:9191"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
)

type ServiceConfig struct {
	DatabaseType    dblayer.DBTYPE `json:"database_type"`
	DBConnection    string         `json:"db_connection"`
	RestfulEndpoint string         `json:"resfulapi_endpoint"`
	RestfulTLSEndpoint string         `json:"resfulapi_tls_endpoint"`
	AMQPMessageBroker string `json:"amqp_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEndpointDefault,
		RestfulTLSEndpointDefault,
		AMQPMessageBrokerDefault,
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)

	if broker := os.Getenv("AMQP_URL"); broker != "" {
		conf.AMQPMessageBroker = broker
	}

	return conf, err
}