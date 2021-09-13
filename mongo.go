package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection(collectionName string) (*mongo.Collection, context.Context, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://recipe-user:123456lyon@cluster0.j2fpf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Recipes").Collection(collectionName)
	return collection, ctx, client
}

func getAllRecipes() []Recipe {
	var recipes []Recipe = []Recipe{}
	collection, ctx, client := getCollection("Recipes")
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	defer client.Disconnect(ctx)
	if err := cur.All(ctx, &recipes); err != nil {
		panic(err)
	}
	return recipes
}

func getRecipe(id int64) Recipe {
	var recipe Recipe = Recipe{}
	collection, ctx, client := getCollection("Recipes")
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&recipe)
	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	return recipe
}

func getRecipeByName(name string) []Recipe {
	var recipes []Recipe = []Recipe{}
	collection, ctx, client := getCollection("Recipes")
	filter := bson.M{"name": bson.M{"$regex": name, "$options": "im"}}
	cur, currErr := collection.Find(ctx, filter)

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	defer client.Disconnect(ctx)
	if err := cur.All(ctx, &recipes); err != nil {
		panic(err)
	}
	return recipes
}

func saveRecipe(recipe Recipe) {
	collection, ctx, client := getCollection("Recipes")
	res, err := collection.InsertOne(ctx, recipe)
	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
func getMaxId() int64 {
	collection, ctx, client := getCollection("Recipes")
	options := options.Find()
	options.SetSort(bson.D{{"id", -1}})
	options.SetLimit(1)
	cur, currErr := collection.Find(ctx, bson.M{}, options)
	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	defer client.Disconnect(ctx)
	var recipes []Recipe = []Recipe{}
	if err := cur.All(ctx, &recipes); err != nil {
		panic(err)
	}
	fmt.Println(recipes)
	return recipes[0].Id
}
