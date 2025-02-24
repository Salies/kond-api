package main

import (
	api "Salies/kond-api/internal/api"
	//	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//	"github.com/joho/godotenv"
)

func main() {
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	api.SteamApiKey = os.Getenv("STEAM_API_KEY")

	api.InitDb()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "we the goat",
		})
	})

	// get players data from Steam
	r.POST("/players", func(c *gin.Context) {
		var playersData api.SteamPlayers
		if err := c.ShouldBindJSON(&playersData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := api.GetPlayersFromSteam(&playersData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	// upload match to kond
	r.POST("/upload", func(c *gin.Context) {
		var matchData api.MatchCreate
		if err := c.ShouldBindJSON(&matchData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := api.InsertMatch(&matchData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	// get match data
	r.GET("/matches/:id", func(c *gin.Context) {
		id := c.Param("id")

		data, err := api.GetMatchById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	// try to retrieve match id by file hash
	r.GET("/demo/:hash", func(c *gin.Context) {
		hash := c.Param("hash")

		data, err := api.GetMatchIdFromFileHash(hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":5000")
}
