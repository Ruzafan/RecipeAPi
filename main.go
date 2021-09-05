package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	log.Fatal(r.Run())
}

func main() {
	fmt.Println("Rest API v1 - Gin Router")
	getRecipes()
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

func returnRecipe(c *gin.Context) {
	fmt.Println("Endpoint Hit: return recipe")
	id := c.Param("id")
	var recipe Recipe = Recipe{}
	for i := 0; i < len(Recipes); i++ {
		if Recipes[i].Id == id {
			recipe = Recipes[i]
			break
		}
	}
	if (recipe == Recipe{}) {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, recipe)
	}
}

func getRecipes() {
	jsonFile, err := os.Open("recipes.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Recipes)
}
