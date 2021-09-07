package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.String(200, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/recipes", returnAllRecipes)
	r.GET("/recipe/:id", returnRecipe)
	r.POST("/recipe/add", setRecipe)
	log.Fatal(r.Run())
}

func main() {
	fmt.Println("Rest API v1 - Gin Router")
	handleRequests()
}

type Recipe struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func returnAllRecipes(c *gin.Context) {
	recipes := getAllRecipes()
	fmt.Println("Endpoint Hit: returnAllRecipes")
	c.JSON(200, recipes)
}

func returnRecipe(c *gin.Context) {
	fmt.Println("Endpoint Hit: return recipe")
	id := c.Param("id")
	recipe := getRecipe(id)
	if (recipe == Recipe{}) {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, recipe)
	}
}

func setRecipe(c *gin.Context) {
	var recipe Recipe = Recipe{}
	recipe.Id = c.PostForm("Id")
	recipe.Name = c.PostForm("Name")
	recipe.Image = c.PostForm("Image")

}
