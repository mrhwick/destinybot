package slackhandlers

import slackbot "github.com/mrhwick/go-slackbot"

// GetTriggerPhraseMapping suggests the mapping of triggers to slack message handlers
func GetTriggerPhraseMapping() map[string]slackbot.MessageHandler {
	return map[string]slackbot.MessageHandler{
		"^ping$": PingHandler,
	}
}
