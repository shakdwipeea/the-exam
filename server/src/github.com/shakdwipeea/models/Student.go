package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/shakdwipeea/utils"
)

const studentCollectionName string = "Student"

type Student struct {
	Username string
	Password string
	Group string
}

func GetUserNames(db *mgo.Database) ([]string, error) {
	var students []Student
	var userNames []string
	err := db.C(studentCollectionName).Find(bson.M{}).All(&students)

	if err != nil {
		return userNames, err
	}

	for _, student := range students {
		userNames = append(userNames, student.Username)
	}

	return userNames, nil
}

func (s *Student) NewUser(db *mgo.Database) error {
	s.Password = utils.HashPassword(s.Password)
	return db.C(studentCollectionName).Insert(s)
}

func (s *Student) CheckUserPassword(db *mgo.Database) (Student, error) {
	s.Password = utils.HashPassword(s.Password)

	var stud Student
	err := db.C(studentCollectionName).Find(bson.M{
		"username": s.Username,
		"password": s.Password,
	}).One(&stud)

	return stud, err
}