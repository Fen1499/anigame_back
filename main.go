package main

import (
	"fmt"
	"net/http"
	"nothing/anigame/controllers"
	"nothing/anigame/pkg"

	"github.com/gin-gonic/gin"
)

type GuessRequest struct {
	Number int32  `json:"number"`
	Letter string `json:"letter"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := pkg.OpenConn()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get daily
	r.GET("/daily", func(c *gin.Context) {
		picked_on, err := controllers.Daily(db)

		c.JSON(http.StatusOK, gin.H{"value": picked_on, "error": err.Error()})
	})

	r.POST("/guess", func(c *gin.Context) {
		var r GuessRequest
		err := c.BindJSON(&r)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{"number": r.Number, "letter": r.Letter, "error": err})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
