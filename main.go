package main

import (
	"flag"
	"net/http"

	"github.com/mrhwick/destinybot/slackhandlers"
	"github.com/mrhwick/destinybot/webhandlers"
	slackbot "github.com/mrhwick/go-slackbot"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var configFile = flag.String("config", "./config.yaml", "Specify path to config file.")
var debug = flag.Bool("debug", false, "Specify whether this bot should be run in debug mode.")

func init() {
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if *debug {
		logger, _ := zap.NewDevelopmentConfig().Build()
		zap.ReplaceGlobals(logger)
	} else {
		logger, _ := zap.NewProductionConfig().Build()
		zap.ReplaceGlobals(logger)
	}

	if err != nil {
		panic(err)
	}
}

func main() {

	slackToken := viper.GetString("slack.token")
	bot := slackbot.New(slackToken)

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()

	// Slack Handler Config
	for triggerPhrase, messageHandler := range slackhandlers.GetTriggerPhraseMapping() {
		toMe.Hear(triggerPhrase).MessageHandler(messageHandler)
	}

	zap.L().Info("Starting destinybot for slack RTM")
	go bot.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", webhandlers.HelloWorldHandler)

	hostPort := "0.0.0.0:3000"

	zap.L().Sugar().Infof("Now serving destinybot API at %s", hostPort)

	server := &http.Server{Addr: hostPort, Handler: mux}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		zap.L().Fatal("Error listening", zap.Error(err))
	}
	zap.L().Info("Server gracefully stopped")

}
