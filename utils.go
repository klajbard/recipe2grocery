package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

func getSimpleButton(text string) *slack.ButtonBlockElement {
	btnText := slack.NewTextBlockObject("plain_text", text, false, false)
	btn := slack.NewButtonBlockElement("", text, btnText)
	return btn
}

func getCheckboxes(options ...*slack.OptionBlockObject) *slack.SectionBlock {
	text := slack.NewTextBlockObject(slack.PlainTextType, "Grocery list", false, false)
	checkboxes := slack.NewCheckboxGroupsBlockElement("grocery", options...)
	accessory := slack.NewAccessory(checkboxes)
	return slack.NewSectionBlock(text, nil, accessory)
}

func SendList(list *ItemList) {
	options := make([]*slack.OptionBlockObject, 0, len(list.List))
	for _, listElement := range list.List {
		checkboxText := slack.NewTextBlockObject(slack.PlainTextType, listElement, false, false)
		checkboxOption := slack.NewOptionBlockObject(listElement, checkboxText, nil)
		options = append(options, checkboxOption)
	}

	checkboxes := getCheckboxes(options...)

	btncancel := getSimpleButton("Cancel")
	actionBlock := slack.NewActionBlock("remove", btncancel)

	_, _, err := SlackBot.PostMessage(os.Getenv("SLACK_CHANNEL"), slack.MsgOptionBlocks(checkboxes, actionBlock), slack.MsgOptionIconEmoji(":white_check_mark:"))
	if err != nil {
		log.Printf("Posting message failed: %v", err)
	}
}
