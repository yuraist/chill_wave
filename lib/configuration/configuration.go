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
)

type ServiceConfig struct {
	DatabaseType    dblayer.DBTYPE `json:"database_type"`
	DBConnection    string         `json:"db_connection"`
	RestfulEndpoint string         `json:"resfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEndpointDefault,
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}