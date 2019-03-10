package db

import (
	"log"
	
	"github.com/globalsign/mgo"
	//"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var (
	Session		*mgo.Session
	Gridfs		*mgo.GridFS
)

func Init() {
	Session, err := mgo.DialWithInfo(GetConfig().DB)
	if err != nil {
		log.Fatal("CreateSession: %s\n", err)
	}	
	
	Gridfs = Session.DB(AuthDatabase).GridFS("fs")
	
}

func Close() {
	Session.Close()
}


