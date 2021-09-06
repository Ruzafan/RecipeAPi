package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	log.Fatal(r.Run(":8080"))
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
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://recipe-user:123456lyon@cluster0.j2fpf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Recipes").Collection("Recipes")
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &Recipes); err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

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
