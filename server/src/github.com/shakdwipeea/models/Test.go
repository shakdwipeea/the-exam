package models
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"
)

const testCollectionName string = "Test"

type Test struct {
	Id          bson.ObjectId "_id"
	Subject     string
	QuestionIds []bson.ObjectId
	Enable      bool
	Group       string
	Name        string
}

//add a test
func (t *Test) AddTest(db *mgo.Database) error {
	t.Id = bson.NewObjectId()

	return db.C(testCollectionName).
	Insert(t)
}

//get all tests
func (t *Test) GetAllTest(db *mgo.Database) ([]Test, error) {
	var tests []Test

	err := db.C(testCollectionName).Find(bson.M{
		"subject": t.Subject,
	}).All(&tests)

	if err != nil {
		return nil, err
	}

	return tests, nil
}

/**
	Get test with id
 */
func (t *Test) GetTest(db *mgo.Database) (Test, error) {
	var test Test
	err := db.C(testCollectionName).Find(bson.M{
		"_id": t.Id,
		"subject": t.Subject,
	}).One(&test)

	if err != nil {
		return test, err
	}

	return test, nil
}

//to enable a test can be used to set other props if required
func (t *Test) SetTestProps(db *mgo.Database) error {

	change := bson.M{
		"$set": bson.M{
			"enable": t.Enable,
			"timestamp": time.Now(),
		},
	}

	return db.C(testCollectionName).Update(bson.M{
		"_id": t.Id,
	}, change)
}