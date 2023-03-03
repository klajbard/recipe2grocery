package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database
var Recipes *mgo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Loaded env variables")
	fmt.Println("Loaded config")
	s, err := mgo.Dial(os.Getenv("MONGODB_CLIENT"))
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(os.Getenv("MONGODB_DB"))
	Recipes = DB.C("recipes")

	log.Println("MongoDB connected")
}
