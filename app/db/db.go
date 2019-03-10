package db

import (
	"log"
	
	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var (
	Session		*mgo.Session
)

func Init() {
	
	Session, err := mgo.DialWithInfo(config.GetConfig().DB)
	if err != nil {
		log.Fatal("CreateSession: %s\n", err)
	}	
	
}

func Close() {
	Session.Close()
}


