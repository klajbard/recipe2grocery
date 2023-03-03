package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

func GetRecipes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, _ := strconv.Atoi(ps.ByName("page"))

	count := 10
	offset := (page - 1) * count

	var recipes []Recipe

	err := Recipes.Find(bson.M{}).Skip(offset).Limit(count).All(&recipes)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func GetRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var recipe Recipe

	err := Recipes.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&recipe)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.URL.Query().Get("id")

	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe.Slug = slug.Make(recipe.Title)
	document := bson.M{"slug": recipe.Slug, "title": recipe.Title, "description": recipe.Description, "ingredients": recipe.Ingredients}

	if id != "" {
		Recipes.Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": document})
	} else {
		Recipes.Insert(document)
	}
}

func SendRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var list ItemList
	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendList(&list)
}

func App(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	index := template.Must(template.ParseFiles("client/dist/index.html"))
	index.Execute(w, nil)
	defer r.Body.Close()
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("client/dist/index.html"))
	index.Execute(w, nil)
	defer r.Body.Close()
}
