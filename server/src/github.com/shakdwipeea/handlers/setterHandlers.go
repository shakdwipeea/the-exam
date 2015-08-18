package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shakdwipeea/models"
	"github.com/shakdwipeea/utils"
	"net/http"
)

func (m *Mongo) LoginSetter(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&req)

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error parsing request.",
		})
		return
	}

	var setter models.QuestionSetter

	setter.Username = req.Username
	setter.Password = req.Password

	questionSetter, err := setter.GetSetter(m.Database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error occured while getting the setter",
		})
		return
	} else if questionSetter.Verified == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "You have not been verified",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["username"] = questionSetter.Username

	tokenString, err := token.SignedString([]byte(utils.MySigningKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error occured while creating a token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err":      false,
		"msg":      "Got",
		"username": questionSetter.Username,
		"token":    tokenString,
	})

}

func (m *Mongo) SignupSetter(c *gin.Context) {
	var req struct {
		Username string `"json:username"`
		Password string `"json:password"`
	}

	err := c.BindJSON(&req)

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Error parsing request.",
		})
		return
	}

	var setter models.QuestionSetter
	setter.Username = req.Username
	setter.Password = req.Password
	setter.Verified = false

	found := setter.GetSetterUserName(m.Database)

	if found == true {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Already exists",
		})
		return
	}

	err = setter.AddSetter(m.Database)

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": true,
			"msg": "Cannot add you",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": false,
		"msg": "You have been added. Pending verification from admin",
	})
}
