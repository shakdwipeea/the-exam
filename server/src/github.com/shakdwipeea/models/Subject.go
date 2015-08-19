package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const subjectCollectionName string = "Subject"

type Subject struct {
	Name string
}

func (s *Subject) Get(db *mgo.Database) []Subject {
	var subject []Subject

	err := db.C(subjectCollectionName).
		Find(bson.M{}).All(&subject)

	if err != nil {
		log.Println("Error", err)
	}

	return subject
}

func (s *Subject) Add(db *mgo.Database) error {
	if s.Name == "" {
		return errors.New("Name not there Why???")
	}

	err := db.C(subjectCollectionName).
		Insert(s)

	return err
}
