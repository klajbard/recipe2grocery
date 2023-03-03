package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

var SlackBot *slack.Client

func initSlack() {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	SlackBot = slack.New(botToken, slack.OptionAppLevelToken(appToken))

	client := socketmode.New(SlackBot)

	client.Run()
}

func init() {
	initSlack()
}

func main() {
	mux := httprouter.New()
	mux.GET("/", App)
	mux.GET("/recipes/:page", GetRecipes)
	mux.POST("/recipe", CreateRecipe)
	mux.GET("/recipe/:id", GetRecipe)
	mux.POST("/recipe/send", SendRecipe)

	mux.ServeFiles("/assets/*filepath", http.Dir("client/dist/assets"))
	// mux.HandleMethodNotAllowed = false
	mux.NotFound = http.HandlerFunc(NotFoundHandler)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8081", handler)
}
