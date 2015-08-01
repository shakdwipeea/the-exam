package models
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const resultCollectionName string = "Result"

type Answer struct {
	QuestionId bson.ObjectId
	Answer     string
}

type Result struct {
	Username string
	Score    string
	TestId   bson.ObjectId
	Response []Answer
}

func (r *Result) InsertResult(db *mgo.Database) error {
	return db.C(resultCollectionName).Insert(r)
}