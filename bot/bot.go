package bot

import (
	"regexp"
	"fmt"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/message"
	"github.com/Knetic/govaluate"
)

type Bot struct {
	Name string
	hub *hub.Hub
}

func CreateBot(name string, hub *hub.Hub) *Bot {
	bot := &Bot{
		Name: name,
		hub: hub,
	}

	hub.Register <- bot

	return bot
}

func (bot *Bot) Send(msg message.Message) {
	// Define the bots hail signature
	hail := regexp.MustCompile("^@" + bot.Name)
	eval := regexp.MustCompile("^@" + bot.Name + "\\s(.+)")

	switch {
		case eval.MatchString(msg.Message):
			str := eval.FindStringSubmatch(msg.Message)[1]
			expr, err := govaluate.NewEvaluableExpression(str)
			if err != nil { bot.fail(); return }

			// If expression is valid, continue
			if res, err := expr.Evaluate(nil); err != nil {
				bot.fail()
			} else {
				bot.send(fmt.Sprintf("%v", res))
			}
		case hail.MatchString(msg.Message):
			bot.send("Hi, " + msg.Username + "! What can I do for you?")
		default:
	}
}

func (bot *Bot) send(msg string) {
	bot.hub.Broadcast <- message.Message{
		Username: bot.Name,
		Message: msg,
	}
}

func (bot *Bot) fail() {
	bot.send("Sorry, I didn't get that :(")
}