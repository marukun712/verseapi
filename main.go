package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Id    string `json:"id"`
	Time  string `json:"time"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://okaramap.netlify.app"},
	}))

	router.GET("/api", api)
	router.GET("/", root)
	router.Run(":5000")
}

func api(c *gin.Context) {
	items := []Item{}
	file, err := os.Open("./api/res.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&items); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, items)
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To VERSEAPI",
	})
}
