package main

type Ingredient struct {
	Name   string  `json:"name" bson:"name"`
	Amount float32 `json:"amount" bson:"amount"`
	Unit   string  `json:"unit" bson:"unit"`
}

type Recipe struct {
	Slug        string       `json:"slug" bson:"slug"`
	Title       string       `json:"title" bson:"title"`
	Description string       `json:"description" bson:"description"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
}

type ItemList struct {
	List []string `json:"list"`
}
