package mongolayer

import (
	"gopkg.in/mgo.v2"
)

const (
	DB = "chill_wave"
	USERS = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)

	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{session: s,}, err
}

