package mgo

import "gopkg.in/mgo.v2"

var MySession *mgo.Session

func init(){
	DB_URL := "mongodb://localhost:27017/PageNote"
	session, err := mgo.Dial( DB_URL )
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	MySession = session
}