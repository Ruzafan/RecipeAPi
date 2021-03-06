package main

import (
	"fmt"
	"log"
	"strconv"

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
	r.GET("/recipe/find/:name", findRecipe)
	r.POST("/recipe/add", setRecipe)
	log.Fatal(r.Run())
}

func main() {
	fmt.Println("Rest API v1 - Gin Router")
	handleRequests()
}

type Recipe struct {
	Id    int64  `json:"id"`
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
	i, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		fmt.Println(err)
	}
	recipe := getRecipe(i)
	if (recipe == Recipe{}) {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, recipe)
	}
}

func findRecipe(c *gin.Context) {
	fmt.Println("Endpoint Hit: return recipe")
	name := c.Param("name")
	recipes := getRecipeByName(name)
	c.JSON(200, recipes)
}

func setRecipe(c *gin.Context) {
	var recipe Recipe = Recipe{}
	c.BindJSON(&recipe)
	recipe.Id = getMaxId() + 1
	saveRecipe(recipe)
	c.JSON(201, recipe)
}
