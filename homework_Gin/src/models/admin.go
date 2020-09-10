package model

import "gopkg.in/mgo.v2/bson"

type Admin struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Login    string        `bson:"login" json:"login"`
	Password string        `bson:"password" json:"password"`
}
