package main

import (
	"net/http"
	"nothing/anigame/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := pkg.OpenConn()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get daily
	r.GET("/daily", func(c *gin.Context) {
		var romanji string
		db.QueryRow("SELECT romanji FROM daily where picked_on = $1", time.Now().Format("2006-01-02")).Scan(&romanji)

        // Would this even return anything? idk because each user will have it's own daily
		c.JSON(http.StatusOK, gin.H{"date": time.Now().Format("2006-01-02"), "value": romanji})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
