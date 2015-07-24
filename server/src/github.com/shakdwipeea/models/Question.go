package models
import "gopkg.in/mgo.v2"

const questionCollectionName string = "Question"

type Question struct {
	Id float64
	QuestionText string
	Option1 string
	Option2 string
	Option3 string
	Option4 string
	Tags []string
	Subject string
}

func (q *Question) AddQuestion (db *mgo.Database) error {
	err := db.C(questionCollectionName).
		Insert(q)

	return err
}