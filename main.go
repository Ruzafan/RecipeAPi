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
	log.Fatal(r.Run())
}

func main() {
	fmt.Println("Rest API v1 - Gin Router")
	Recipes = []Recipe{
		Recipe{Id: "1", Name: "Oreo cupcake", Image: "https://www.recetin.com/wp-content/uploads/2014/01/cupcakes__oreo.jpg"},
	}
	handleRequests()
}

type Recipe struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

var Recipes []Recipe

func returnAllRecipes(c *gin.Context) {
	fmt.Println("Endpoint Hit: returnAllRecipes")
	c.JSON(200, Recipes)
}
