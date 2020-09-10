package model

import "gopkg.in/mgo.v2/bson"

type Task struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string
	State bool
}

type Tasks []Task
