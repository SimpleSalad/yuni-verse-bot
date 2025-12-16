package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

//Variables for parameters
var (
	BotToken string
)

//Define the bot token
func init() {
	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.Parse()
}

//
func main() {
	
	dg, err := discordgo.New("Bot " + BotToken)
	
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
	}
}