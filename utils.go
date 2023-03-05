package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/slack-go/slack"
)

func getSimpleButton(text string) *slack.ButtonBlockElement {
	btnText := slack.NewTextBlockObject("plain_text", text, false, false)
	btn := slack.NewButtonBlockElement("", text, btnText)
	return btn
}

func getCheckboxes(title string, options ...*slack.OptionBlockObject) *slack.SectionBlock {
	text := slack.NewTextBlockObject(slack.PlainTextType, title, false, false)
	checkboxes := slack.NewCheckboxGroupsBlockElement("grocery", options...)
	accessory := slack.NewAccessory(checkboxes)
	return slack.NewSectionBlock(text, nil, accessory)
}

func getMaxValue(value int, value2 int, max int) {

}

func SendList(list *ItemList) (err error) {
	optionsLen := len(list.List)
	fullIteration := optionsLen / 10
	if optionsLen%10 > 0 {
		fullIteration += 1
	}
	iteration := 0

	// Slack allows only 10 checkboxes in a single message block therefor
	// we are sending multiple messages
	for fullIteration > iteration {
		var endIndex int
		if fullIteration > iteration+1 {
			endIndex = (iteration + 1) * 10
		} else {
			endIndex = iteration*10 + optionsLen%10
		}
		options := make([]*slack.OptionBlockObject, 0, 10)

		for _, listElement := range list.List[iteration*10 : endIndex] {
			checkboxText := slack.NewTextBlockObject(slack.PlainTextType, listElement, false, false)
			checkboxOption := slack.NewOptionBlockObject(slug.Make(listElement), checkboxText, nil)
			options = append(options, checkboxOption)
		}

		checkboxes := getCheckboxes(fmt.Sprintf("Grocery list #%d", iteration+1), options...)

		btncancel := getSimpleButton("Cancel")
		actionBlock := slack.NewActionBlock("remove", btncancel)

		_, _, err = SlackBot.PostMessage(os.Getenv("SLACK_CHANNEL"), slack.MsgOptionBlocks(checkboxes, actionBlock), slack.MsgOptionIconEmoji(":white_check_mark:"))
		if err != nil {
			log.Printf("Posting message failed: %v", err)
		}
		iteration += 1
	}

	return
}
