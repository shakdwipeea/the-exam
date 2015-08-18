package models

import (
	"github.com/shakdwipeea/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const questionSetterCollectionName string = "QuestionSetter"

type QuestionSetter struct {
	Username string
	Password string
	Verified bool
}

func (setter *QuestionSetter) GetSetter(db *mgo.Database) (QuestionSetter, error) {
	setter.Password = utils.HashPassword(setter.Password)

	var QSetter QuestionSetter
	err := db.C(questionSetterCollectionName).Find(bson.M{
		"username": setter.Username,
		"password": setter.Password,
	}).One(&QSetter)

	return QSetter, err
}

func (setter *QuestionSetter) AddSetter(db *mgo.Database) error {
	setter.Password = utils.HashPassword(setter.Password)

	return db.C(questionSetterCollectionName).Insert(setter)
}

func (setter *QuestionSetter) GetSetterUserName(db *mgo.Database) bool {
	var set QuestionSetter
	err := db.C(questionSetterCollectionName).Find(bson.M{
		"username": setter.Username,
	}).One(&set)

	if err != nil {
		return false
	} else {
		return true
	}
}
