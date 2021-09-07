package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection(collectionName string) (*mongo.Collection, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://recipe-user:123456lyon@cluster0.j2fpf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("Recipes").Collection(collectionName)
	return collection, ctx
}

func getAllRecipes() []Recipe {
	var recipes []Recipe = []Recipe{}
	collection, ctx := getCollection("Recipes")
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	if err := cur.All(ctx, &recipes); err != nil {
		panic(err)
	}
	return recipes
}

func getRecipe(id string) Recipe {
	var recipe Recipe = Recipe{}
	collection, ctx := getCollection("Recipes")
	err := collection.FindOne(ctx, bson.M{"Id": id}).Decode(&recipe)
	if err != nil {
		panic(err)
	}
	return recipe
}
