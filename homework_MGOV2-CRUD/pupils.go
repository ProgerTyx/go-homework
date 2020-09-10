package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Pupil struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
	Rating    float64       `bson:"rating" json:"rating"`
}

type Pupils []Pupil

func getPupils() Pupils {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("School-journal").C("Pupils")

	var pupils []Pupil
	err = c.Find(nil).All(&pupils)

	return pupils
}

func getPupilById(id string) Pupil {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return Pupil{}
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("School-journal").C("Pupils")
	var pupil Pupil
	err = c.FindId(bson.ObjectIdHex(id)).One(&pupil)

	return pupil
}

func (p *Pupil) createPupil() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("School-journal").C("Pupils")
	err = c.Insert(p)
	if err != nil {
		log.Println("Error creating Pupil: ", err.Error())
		return false
	}
	return true
}

func deletePupil(id string) bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("School-journal").C("Pupils")
	err = c.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println("Pupil not found: ", err.Error())
		return false
	}

	return true
}
