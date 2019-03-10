package db

import (
	"time"
	"github.com/globalsign/mgo"
)

const (
	MongoDBHosts = "127.0.0.1:27017"
	AuthDatabase = "todoDB"
	AuthUserName = "todoOwner"
	AuthPassword = "todomanager"
	COLLECTION 	 = "todoC"
)

type Config struct {
	DB	*mgo.DialInfo
}

func GetConfig() *Config {
	return &Config {
		DB: &mgo.DialInfo {
			Addrs:		[]string{MongoDBHosts},
			Timeout:	60 * time.Second,
			Database:	AuthDatabase,
			Username:	AuthUserName,
			Password:	AuthPassword,
		},
	}
}

