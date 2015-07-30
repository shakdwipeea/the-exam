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
	"gopkg.in/mgo.v2/bson"
)



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
			"msg": "Error parsing request.",
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

		tokenString, err := token.SignedString([]byte(utils.MySigningKey))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": true,
				"msg": "Error occured while creating a token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"err":     false,
			"msg":     "Got",
			"teacher": teacher,
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Not found.",
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

		token, err := utils.JwtAuthenticator(QuestionInput.Token)

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

	token, err1 := utils.JwtAuthenticator(tagInput.Token)

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

/**
	handlers to retrieve questions
	from database
 */
func (m *Mongo) GetQuestions(c *gin.Context) {

	tokenForm := c.Query("token")

	if tokenForm == "" {
		log.Println("HUW", tokenForm)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Parse error",
		})
		return
	}

	token, err := utils.JwtAuthenticator(tokenForm)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Log In Again",
		})
		return
	}

	subject, ok := token.Claims["subject"].(string)

	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Log In Again",
		})
		return
	}

	questions, err := new(models.Question).GetQuestionsOfSubject(m.Database, subject)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": nil,
		"questions": questions,
	})
}

func (m *Mongo) AddTest(c *gin.Context) {
	var addTestInput struct {
		Token string `json:"token"`
		Ids   []string `json:"ids"`
		Name  string `json:"name"`
		Group string `json:"group"`
	}

	//Parse the req body
	err := c.BindJSON(&addTestInput)

	if err != nil {
		//send the incorrect response
		utils.ErrorResponse(c, http.StatusBadRequest, "Parse error")
		return
	}

	//check the token
	token, err := utils.JwtAuthenticator(addTestInput.Token)

	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, "Log In again")
		return;
	}

	var test models.Test

	subject, ok := token.Claims["subject"].(string)

	if !ok {
		utils.ErrorResponse(c, http.StatusForbidden, "Log In again")
		return;
	}

	var questionIds []bson.ObjectId
	//convert string ids to object ids

	for _, e := range addTestInput.Ids {
		if (bson.IsObjectIdHex(e)) {
			log.Println("The objectid for")
			questionIds = append(questionIds, bson.ObjectIdHex(e))
		} else {
			utils.ErrorResponse(c, http.StatusInternalServerError, "ITS NOT U ITS ME ")
			return
		}

	}

	test.Subject = subject
	test.QuestionIds = questionIds
	test.Name = addTestInput.Name
	test.Group = addTestInput.Group
	test.Enable = false

	//add test to database
	err = test.AddTest(m.Database)

	if err != nil {
		log.Println("The funckin eeror", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Cannot Insert")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": nil,
	})

}

func (m *Mongo) GetAllTest(c *gin.Context) {
	tokenReceived := c.Query("token")

	subject, err := utils.AuthenticateTokenGetSubject(tokenReceived)

	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, "Log In Again")
		return
	}


	var test models.Test
	test.Subject = subject

	tests, err := test.GetAllTest(m.Database)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Could not retreive test")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": nil,
		"tests": tests,
	})
}

func (m *Mongo) GetTest(c *gin.Context) {
	tokenReceived := c.Query("token")
	id := c.Param("id")

	subject, err := utils.AuthenticateTokenGetSubject(tokenReceived)

	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, "Log in again")
		return
	}
	log.Println("thr dod", id)
	var test models.Test
	test.Subject = subject
	if bson.IsObjectIdHex(id) {
		test.Id = bson.ObjectIdHex(id)
	} else {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Oh boy not again")
		return
	}

	test, err = test.GetTest(m.Database)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Could not get the test")
		return
	}

	/**
	todo get all questions for this test
	should be probably done with go routines for
	improved performance
	 */
	var questions []models.Question
	for _, id := range test.QuestionIds {
		temp_question, err := getQuestion(m.Database, id)

		if err != nil {
			log.Println("Somethin's up ", err)
		} else {
			questions = append(questions, temp_question)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"err": nil,
		"test": test,
		"questions": questions,
	})
}

func getQuestion(db *mgo.Database, questionId bson.ObjectId) (models.Question, error) {
	question, err := new(models.Question).GetQuestion(db, questionId)

	if err != nil {
		return question, err
	}

	return question, nil
}




