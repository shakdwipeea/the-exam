package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const resultCollectionName string = "Result"

//Answer Struct for holding response
type Answer struct {
	QuestionId bson.ObjectId
	Answer     string
}

//Result struct to hold the result of a test
type Result struct {
	Username string
	Score    string
	TestId   bson.ObjectId
	Response []Answer
}

//InsertResult func to insert the result passed
func (r *Result) InsertResult(db *mgo.Database) error {
	return db.C(resultCollectionName).Insert(r)
}

//GetResults func to get all the test results
func GetResults(db *mgo.Database) ([]Result, error) {
	var results []Result
	err := db.C(resultCollectionName).Find(bson.M{}).All(&results)

	return results, err
}

//GetResultByID func to get results of given test id
func (r *Result) GetResultByID(db *mgo.Database) ([]Result, error) {
	var results []Result
	err := db.C(resultCollectionName).Find(bson.M{
		"testid": r.TestId,
	}).Select(bson.M{
		"testid":   1,
		"score":    1,
		"username": 1,
	}).All(&results)

	return results, err
}
