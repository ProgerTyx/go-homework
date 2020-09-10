package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
}

type Users []User
