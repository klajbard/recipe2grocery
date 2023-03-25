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

func sendMessage(w http.ResponseWriter, message string, errorCode ...int) {
	resp := ApiResponse{
		Message: message,
		Success: true,
	}
	w.Header().Set("Content-Type", "application/json")

	if len(errorCode) > 0 && errorCode[0] > 0 {
		w.WriteHeader(errorCode[0])
		resp.Success = false
	}

	json.NewEncoder(w).Encode(resp)
}

func sendData(w http.ResponseWriter, data interface{}) {
	resp := ApiResponse{
		Data:    data,
		Success: true,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func GetRecipes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, _ := strconv.Atoi(ps.ByName("page"))

	count := 10
	offset := (page - 1) * count

	var recipes []Recipe

	err := Recipes.Find(bson.M{}).Skip(offset).Limit(count).All(&recipes)

	if err != nil {
		fmt.Println(err)
		sendMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendData(w, recipes)
}

func GetRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var recipe Recipe

	err := Recipes.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&recipe)

	if err != nil {
		fmt.Println(err)
		sendMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendData(w, recipe)
}

func RemoveRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	err := Recipes.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		fmt.Println(err)
		sendMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, fmt.Sprintf("Successfully removed recipe with id: %s", id))
}

func UpsertRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.URL.Query().Get("id")

	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)

	if err != nil {
		sendMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe.Slug = slug.Make(recipe.Title)
	document := bson.M{"slug": recipe.Slug, "title": recipe.Title, "description": recipe.Description, "ingredients": recipe.Ingredients}

	if id != "" {
		err = Recipes.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": document})
	} else {
		err = Recipes.Insert(document)
	}

	if err != nil {
		sendMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Successfully updated recipe")
}

func SendRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var list ItemList
	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		sendMessage(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = SendList(&list)

	if err != nil {
		sendMessage(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Successfully send grocery list to Slack")
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
