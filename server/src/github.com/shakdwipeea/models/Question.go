package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const questionCollectionName string = "Question"

type Question struct {
	Id           bson.ObjectId "_id"
	QuestionText string
	Option1      string
	Option2      string
	Option3      string
	Option4      string
	Correct      string
	Tags         []string
	Subject      string
}

func (q *Question) AddQuestion (db *mgo.Database) error {
	/*	q.Id = new(bson.ObjectId)*/

	err := db.C(questionCollectionName).
		Insert(q)

	return err
}

func (q *Question) GetQuestionsOfSubject(db *mgo.Database, subject string) ([]Question, error) {
	var questions []Question

	err := db.C(questionCollectionName).
	Find(bson.M{
		"subject": subject,
	}).
	All(&questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}