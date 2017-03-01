package models
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/shakdwipeea/utils"
)

const teacherCollectionName string = "Teacher"

/**
Mongo docs
 */

type Teacher struct {
	Username string
	Password string
	Subject string
}

func (t *Teacher) AddDoc (db *mgo.Database) utils.Response  {

	/**
	Hash the password
	 */
	t.Password = utils.HashPassword(t.Password)

	var teachers []Teacher
	/*
	Dangerous way improvemtn required
	 */
	err := db.C(teacherCollectionName).Find(bson.M{
		"username": t.Username,
	}).All(&teachers)

	if err != nil || len(teachers) > 0 {
		return utils.Response{true, "Already exist"}
	} else {
		err1 := db.C(teacherCollectionName).Insert(t)

		if err1 != nil {
			return utils.Response{true, "Error ro"}
		} else {
			return utils.Response{false, "Inserted"}
		}
	}
}

func (t *Teacher) GetByUsernameAndPassword (db *mgo.Database) Teacher {

	t.Password = utils.HashPassword(t.Password)

	var tempTeacher  Teacher
	err := db.C(teacherCollectionName).
		Find(bson.M{
			"username": t.Username,
			"password": t.Password,
		}).
		One(&tempTeacher)

	if err != nil {
		log.Print("Coming", t.Username, t.Password);
		log.Print("Error in query", err)
		return tempTeacher
	} else {
		return tempTeacher
	}
}
