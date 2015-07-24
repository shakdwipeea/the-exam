package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shakdwipeea/models"
	"github.com/shakdwipeea/utils"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"
	"errors"
)

const mySigningKey string = "madh165r35##@@#"

/**
The database struct which contains all the handlers
*/

type Mongo struct {
	Database *mgo.Database
}

/**
Sample type to represent mongo docs
*/

type Location struct {
	Hid      string
	Type     string
	Lat      float64
	Lng      float64
	Name     string
	Place_id string
}

/*
For index route
*/
func (mongo *Mongo) DetailHandler(c *gin.Context) {
	var res struct {
		Err     bool
		Message string
	}

	res.Err = true
	res.Message = "Hi"

	result := []Location{}

	restaurants := mongo.Database.C("location")
	err := restaurants.Find(nil).All(&result)

	for _, value := range result {
		res.Message += value.Name
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
	}

	c.JSON(http.StatusOK, res)
}

/**
Handler for new teacher
*/
func (mongo *Mongo) AddTeacher(c *gin.Context) {
	/**
	Room for imporv use auto bindind
	*/

	var teacher models.Teacher

	adminPass := c.PostForm("adminPass")

	if adminPass == "admin" {

		teacher.Password = c.PostForm("password")
		teacher.Subject = c.PostForm("subject")
		teacher.Username = c.PostForm("username")

		if teacher.Password == "" || teacher.Subject == "" || teacher.Username == "" {
			var res utils.Response
			res.Err = true
			res.Msg = "Missing params"

			c.JSON(http.StatusForbidden, res)
		} else {

			/**
			Insert the teacher in the db
			*/
			response := teacher.AddDoc(mongo.Database)

			if response.Err == true {
				c.JSON(http.StatusInternalServerError, response)
			} else {
				c.JSON(http.StatusOK, response)
			}
		}

	} else {
		var res utils.Response
		res.Err = true
		res.Msg = "Unauth "

		c.JSON(http.StatusForbidden, res)
	}
}

func (mongo *Mongo) Login(c *gin.Context) {

	var Incoming struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&Incoming)

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error parsing reques",
		})
		return
	}

	var t models.Teacher

	t.Username = Incoming.Username
	t.Password = Incoming.Password

	println("Got from user", Incoming.Username, t.Username, t.Password)

	teacher := t.GetByUsernameAndPassword(mongo.Database)

	if teacher.Subject != "" {
		/**
		Create the token
		*/
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims["username"] = teacher.Username
		token.Claims["subject"] = teacher.Subject
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString([]byte(mySigningKey))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": true,
				"msg": "Error occured while creating a token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"err":     false,
			"msg":     "Got ",
			"teacher": teacher,
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Not found",
		})
	}
}

func (mongo *Mongo) AddQuestion(c *gin.Context) {
	var QuestionInput struct {
		Text    string `json:"questionText"`
		Option1 string `json:"option1"`
		Option2 string `json:"option2"`
		Option3 string `json:"option3"`
		Option4 string `json:"option4"`
		Token   string `json:"token"`
		Subject string `json:"subject"`
		Tags []string `json:"tags"`
		Correct string `json:"correct"`
	}

	err := c.BindJSON(&QuestionInput)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error parsing",
		})
		return
	}

	if QuestionInput.Text != "" && QuestionInput.Option1 != "" &&
		QuestionInput.Option2 != "" && QuestionInput.Option3 != "" &&
		QuestionInput.Option4 != "" && QuestionInput.Subject != "" &&
		QuestionInput.Token != "" && QuestionInput.Correct != "" {

		/*
			validate jwt token
		*/
/*		token, err := jwt.Parse(QuestionInput.Token, func(token *jwt.Token) (interface{}, error) {
			if token.Claims["subject"] == QuestionInput.Subject {
				return []byte(mySigningKey), nil
			} else {
				return nil, nil
			}

		})*/

		token,err := jwtAuthenticator(QuestionInput.Token)

		if err == nil && token.Valid &&
			token.Claims["subject"] == QuestionInput.Subject  {
			/**
			Token is gud . Now move
			*/
			var question models.Question
			question.QuestionText = QuestionInput.Text
			question.Option1 = QuestionInput.Option1
			question.Option2 = QuestionInput.Option2
			question.Option3 = QuestionInput.Option3
			question.Option4 = QuestionInput.Option4
			question.Subject = QuestionInput.Subject
			question.Tags = QuestionInput.Tags
			question.Correct = QuestionInput.Correct

			err := question.AddQuestion(mongo.Database)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": true,
					"msg": "No insert",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"err": false,
				"msg": "Quesiton Inserted",
			})
		} else {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": true,
				"msg": "Sth fishy Login Again",
			})

		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": true,
			"msg": "Missing Params",
		})
	}
}

func (m *Mongo) GetTags(c *gin.Context) {
	var tag []models.Tags

	log.Println(c.ClientIP())

	tag = new(models.Tags).Get(m.Database)
	c.JSON(http.StatusOK, gin.H{
		"err":  false,
		"tags": tag,
	})
}

func (mongo *Mongo) AddTag(c *gin.Context) {
	var tagInput struct {
		Name string `json:"name"`
		Token string `json:"token"`
	}

	err := c.BindJSON(&tagInput)

	var tag models.Tags
	tag.Name = tagInput.Name

	log.Println("Why u do this",tagInput.Name)

	if err != nil {
		log.Println("Add tags err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error Parsing",
		})
		return
	}

	token,err1 := jwtAuthenticator(tagInput.Token)

	if err1 != nil || !token.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Login Again",
		})
		return
	}

	err = tag.Add(mongo.Database)

	if err != nil {
		log.Println("Add tags db err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error Inserting",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": false,
		"msg": "Tag Added",
	})

}

func jwtAuthenticator (token string) (*jwt.Token,error) {
	parsedToken, err := jwt.Parse(token, func (token *jwt.Token) (interface {}, error) {
/*		if _,ok := token.Method.(*jwt.SigningMethodHS256); !ok {
			return nil, errors.New("Wrong token")
		}*/
		return []byte(mySigningKey), nil
	})

	if err != nil && !parsedToken.Valid {
		return nil,errors.New("Tampered Token")
	}

	return parsedToken, nil
}
