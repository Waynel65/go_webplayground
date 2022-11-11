package main

import (
	"net/http"
	"time"
	"github.com/rs/xid"

	"github.com/gin-gonic/gin"
)

/*
	backtick notation:
	- e.g. `json:"name"`
	- This allows us to map each field to a different name when we send them as responses,
	since JSON and Go have different naming conventions.
*/

type Recipe struct{
	ID			 string    `json: "id"`
	Name		 string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe // a global variable that stores the list of recipes
func init(){
	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c * gin.Context){
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String() // xid generates a random id
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe) // structure the response as json
}


func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler) // direct post request towards route /recipes to NewRecipeHandler
	router.Run("localhost:8080")
}