package handlers

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shakdwipeea/models"
	"github.com/shakdwipeea/utils"
	"gopkg.in/mgo.v2/bson"
)

func (m *Mongo) GetUserNames(c *gin.Context) {
	usernames, err := models.GetUserNames(m.Database)

	if err != nil {
		println("E", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Fuck's gone wrong")
		return
	}

	c.JSON(http.StatusOK, usernames)
}

func (m *Mongo) StudentLogin(c *gin.Context) {
	var loginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&loginForm)

	if err != nil {
		log.Println("Github retards", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Incorrect params")
		return
	}

	var studMahan models.Student
	studMahan.Username = loginForm.Username
	studMahan.Password = loginForm.Password

	stud, err := studMahan.CheckUserPassword(m.Database)

	if err != nil {
		log.Println("Absent", err)
		utils.ErrorResponse(c, http.StatusNotFound, "No such user")
		return
	}

	if stud.Username != "" {
		//make the fuckin token

		/**
		Create the token
		*/
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims["username"] = stud.Username
		token.Claims["group"] = stud.Group
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString([]byte(utils.MySigningKey))

		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Error creating token")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	} else {
		utils.ErrorResponse(c, http.StatusNotFound, "No such user")
		return
	}
}

func (m *Mongo) StudentSignUp(c *gin.Context) {
	var signUpForm struct {
		Username string   `json:"username"`
		Password string   `json:"password"`
		Group    []string `json:"groups"`
	}

	err := c.BindJSON(&signUpForm)

	if err != nil {
		log.Println("rui", err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Params wring")
		return
	}

	if signUpForm.Username == "" || signUpForm.Password == "" || len(signUpForm.Group) == 0 {
		log.Println(signUpForm)
		utils.ErrorResponse(c, http.StatusBadRequest, "Give Proper Formats")
		return
	}

	student := models.Student{
		Username: signUpForm.Username,
		Password: signUpForm.Password,
		Group:    signUpForm.Group,
	}

	err = student.NewUser(m.Database)

	if err != nil {
		log.Println("kauwa", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Cannot sign u up")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "All done",
	})
}

func (m *Mongo) GetExams(c *gin.Context) {
	tokenReceived := c.Query("token")
	token, err := utils.JwtAuthenticator(tokenReceived)

	println("ofijwejio")
	if err != nil {
		log.Println("kauwa", err)
		utils.ErrorResponse(c, http.StatusForbidden, "Log IN again")
		return
	}

	username, ok := token.Claims["username"].(string)

	if !ok {
		log.Println("kauwa", ok)
		utils.ErrorResponse(c, http.StatusForbidden, "log in afain")
		return
	}

	var s models.Student
	s.Username = username
	groups, err := s.GetGroup(m.Database)

	if err != nil {
		log.Println("kauwa", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Its the fuckif se")
		return
	}

	var tests []models.Test
	for _, group := range groups {
		t, err := models.GetEnabledByGroup(m.Database, group)

		if err == nil {
			for _, value := range t {
				tests = append(tests, value)
			}

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"tests": tests,
	})

}

func (m *Mongo) StoreResult(c *gin.Context) {
	type ResponseForm struct {
		Id     string `json:"id"`
		Answer string `json:"answer"`
	}

	var resultForm struct {
		Token    string         `json:"token"`
		Score    string         `json:"score"`
		Response []ResponseForm `json:"response"`
		TestId   string         `json:"test_id"`
	}

	err := c.BindJSON(&resultForm)

	if err != nil {
		log.Println("rui", err)
		utils.ErrorResponse(c, http.StatusBadRequest, "Params wring")
		return
	}

	token, err := utils.JwtAuthenticator(resultForm.Token)

	println("ofijwejio")
	if err != nil {
		log.Println("kauwa", err)
		utils.ErrorResponse(c, http.StatusForbidden, "Log IN again")
		return
	}

	username, ok := token.Claims["username"].(string)

	if !ok {
		log.Println("kauwa", ok)
		utils.ErrorResponse(c, http.StatusForbidden, "log in afain")
		return
	}

	var res []models.Answer

	for _, r := range resultForm.Response {
		var temp models.Answer
		temp.Answer = r.Answer

		if bson.IsObjectIdHex(r.Id) {
			temp.QuestionId = bson.ObjectIdHex(r.Id)
		}

		res = append(res, temp)
	}

	var r models.Result
	r.Username = username
	r.Response = res
	r.Score = resultForm.Score

	if bson.IsObjectIdHex(resultForm.TestId) {
		r.TestId = bson.ObjectIdHex(resultForm.TestId)
	}

	err = r.InsertResult(m.Database)

	if err != nil {
		log.Println("kauwa", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "Cannot save result")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Done",
	})

}

//ByScore sorting interface
type ByScore []models.Result

func (s ByScore) Len() int { return len(s) }

func (s ByScore) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByScore) Less(i, j int) bool { return s[i].Score < s[j].Score }

//GetLeaderBoardsOfTest handler func to get the leaderboards of a particilar test
func (m *Mongo) GetLeaderBoardsOfTest(c *gin.Context) {
	testId := c.Param("testId")

	if testId == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "No id provided")
		return
	}

	var result models.Result

	//convert testid string to bson.ObjectId
	if bson.IsObjectIdHex(testId) {
		result.TestId = bson.ObjectIdHex(testId)
	} else {
		utils.ErrorResponse(c, http.StatusExpectationFailed, "Cannot retreive test id")
		return
	}

	results, err := result.GetResultByID(m.Database)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Cannot retreive from db")
		return
	}

	sort.Sort(ByScore(results))

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}
