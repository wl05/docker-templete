package db

import (
	"github.com/globalsign/mgo"
	"server/config"
	"fmt"
)

var (
	DB *mgo.Database
	// Session stores mongo session
	Session *mgo.Session
)

func Connect() {
	session, err := mgo.Dial(config.MongoDBUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	Session = session
	DB = session.DB(config.Database)

}

func Close() {
	Session.Close()
}