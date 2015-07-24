package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shakdwipeea/handlers"
	"github.com/tommy351/gin-cors"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	/**
	Connect to database
	*/
	session, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		println("Error occured", err)
		panic("Error in connecting")
	}

	defer session.Close()

	db := session.DB("Question")
	mongo := handlers.Mongo{db}
	println("COnnected to mongodb")

	/**
	Initialize the router
	*/

	router := gin.Default()

	/**
	For debug purpose allowed cors
	*/
	router.Use(cors.Middleware(cors.Options{}))

	//router.LoadHTMLGlob("views/*")

	/**
	html routes
	*/

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is not what you're looing for. Is it???")
	})

	/**
	API Routes
	*/
	router.GET("/tags", mongo.GetTags)
	router.POST("/add_teacher", mongo.AddTeacher)
	router.POST("/login", mongo.Login)
	/**
	Secure routes
	*/
	secure := router.Group("/secure")
	{
		secure.POST("/add_question", mongo.AddQuestion)
		secure.POST("/tags", mongo.AddTag)
	}

	/**
	for static files
	*/
	router.Static("/public", "./app/")

	/**
	Run the server
	*/

	router.Run(":3000")
}
