package models
import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

const tagsCollectionName string = "Tags"

type Tags struct {
	Name string
}

func (t *Tags) Get (db *mgo.Database) []Tags {
	var tags []Tags

	err := db.C(tagsCollectionName).
		Find(bson.M{}).All(&tags)

	if err != nil {
		log.Println("Error", err)
	}

	return tags
}

func (t *Tags) Add (db *mgo.Database) error {
	if t.Name == "" {
		return errors.New("Name not there Why???")
	}

	err := db.C(tagsCollectionName).
				Insert(t)

	return err
}