package gopher

import (
	"regexp"
	"fmt"
	"time"
	"github.com/timojarv/gochat/hub"
	"github.com/timojarv/gochat/message"
	"github.com/Knetic/govaluate"
)

const name string = "gopher"

type Bot struct {
	hub *hub.Hub
}

func Register(hub *hub.Hub) *Bot {
	bot := &Bot{hub}

	hub.Register <- bot

	return bot
}

func (bot *Bot) Send(msg message.Message) {
	// Define the bots hail signature
	hail := regexp.MustCompile("^@" + name)
	eval := regexp.MustCompile("^@" + name + "\\s(.+)")
	mention := regexp.MustCompile("@" + name)

	time.Sleep(500 * time.Millisecond)

	switch {
		case eval.MatchString(msg.Body):
			str := eval.FindStringSubmatch(msg.Body)[1]
			expr, err := govaluate.NewEvaluableExpression(str)
			if err != nil { bot.fail(); return }

			// If expression is valid, continue
			if res, err := expr.Evaluate(nil); err != nil {
				bot.fail()
			} else {
				bot.send(fmt.Sprintf("%v", res))
			}
		case hail.MatchString(msg.Body):
			bot.send("Hi, " + msg.Sender + "! What can I do for you?")
		case mention.MatchString(msg.Body):
			bot.send("I'm here! :D")
		default:
	}
}

func (bot *Bot) send(msg string) {
	bot.hub.Broadcast <- message.New(name, msg)
}

func (bot *Bot) fail() {
	bot.send("Sorry, I didn't get that :(")
}