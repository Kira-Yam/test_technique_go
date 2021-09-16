package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type brainee struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
	Brand  string `json:"brand"`
}

var brainees = []brainee{
	{ID: "42", Text: "Et si on pouvait se faire livrer à manger dans le train ? Prochaine arrêt : livraison", Author: "Lambert", Brand: "Sncf"},
	{ID: "43", Text: "Et si on allait manger gras ?", Author: "Bernard", Brand: "McDonald's"},
}

func getBrainees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, brainees)
}

func getBraineesByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range brainees {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "brainee not found"})
}

func postBrainees(c *gin.Context) {
	var newBrainee brainee

	if err := c.BindJSON(&newBrainee); err != nil {
		return
	}

	brainees = append(brainees, newBrainee)
	c.IndentedJSON(http.StatusCreated, newBrainee)
}

func main() {
	router := gin.Default()
	router.GET("/brainees", getBrainees)
	router.GET("/brainees/:id", getBraineesByID)
	router.POST("brainees", postBrainees)

	router.Run("localhost:8080")
}
