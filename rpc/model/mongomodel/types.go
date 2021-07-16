package mongomodel

import "github.com/globalsign/mgo/bson"

var mongourl = "mongodb://admin:123456@192.168.217.129:27018/"


type User struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

type Student struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	PhoneNumber string        `bson:"phoneNumber"`
}
