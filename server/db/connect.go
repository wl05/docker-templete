package db

import (
	"github.com/globalsign/mgo"
	"fmt"
	"github.com/spf13/viper"
)

var (
	DB *mgo.Database
	// Session stores mongo session
	Session *mgo.Session
)

func Connect() {
	fmt.Println("========MongoDBUrl:",viper.GetString("MongoDBUrl"), "==========")
	fmt.Println("========Database:",viper.GetString("Database"), "==========")
	session, err := mgo.Dial(viper.GetString("MongoDBUrl"))
	if err != nil {
		fmt.Println(err.Error())
	}
	Session = session
	DB = session.DB(viper.GetString("Database"))
}

func Close() {
	Session.Close()
}