package main

import (
	"github.com/gin-gonic/gin"
	"time"
)


/* 
	backtick notation:
	- e.g. `json:"name"`
	- This allows us to map each field to a different name when we send them as responses, 
	since JSON and Go have different naming conventions.
*/

type Recipe struct{
	Name		 string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}


func main() {
	router := gin.Default()
	router.Run()
}