package main

import (
	"log"
	"net/http"

	"example.com/gin-web/db"
	"github.com/gin-gonic/gin"
)

var (
	router     *gin.Engine
	ListenAddr = "0.0.0.0:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Fail to connect to redis: %s", err.Error())
	}
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		user, err := database.GetAllUser()
		if err != nil {
			if err == db.ErrNil {
				c.JSON(http.StatusNotFound, gin.H{"error": "No record found for "})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// fmt.Println(user.Users[0].Username)
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":   "Home Pageeeeee",
				"UserAll": user.Users,
				// "payload": articles,
			},
		)
	})

	router.POST("/points", func(c *gin.Context) {

		var userJson db.User
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.SaveUser(&userJson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// fmt.Println(UserAll)
	})
	router.Run(ListenAddr)
}
