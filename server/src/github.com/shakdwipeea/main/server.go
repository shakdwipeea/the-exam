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

	router.LoadHTMLGlob("./app/*.html")

	/**
	html routes
	*/

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
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
		secure.POST("/question", mongo.AddQuestion)
		secure.POST("/tags", mongo.AddTag)
		secure.GET("/question", mongo.GetQuestions)

		secure.POST("/test", mongo.AddTest)
		secure.GET("/test", mongo.GetAllTest)
		secure.GET("/test/:id", mongo.GetTest)

		secure.GET("/enable/:id", mongo.EnableTest)
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
