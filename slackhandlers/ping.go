package slackhandlers

import (
	slackbot "github.com/beepboophq/go-slackbot"
	"github.com/nlopes/slack"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

//PingHandler is obviously a ping/pong handler...
func PingHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	zap.L().Debug("Received a `ping` command, responding with `pong`")
	bot.Reply(evt, "pong", slackbot.WithTyping)
}
