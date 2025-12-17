package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables for parameters
var (
	BotToken string
)

// Define the bot token
func init() {
	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	//Create bot session
	dg, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageTest)

	//Declare intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//Open session
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//Wait for termination
	fmt.Println("Yuni Bot connected to Discord.")

	dg.UpdateGameStatus(0, "Dance Dance Revolution")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigchan

	dg.Close()
}

// Check every message if it contains trigger words and respond
func messageTest(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Bot ignores messages created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "1" {
		s.ChannelMessageSend(m.ChannelID, "1")
		s.AddHandlerOnce(messageChainTest1)
	}

	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "reply")
	}
}

func messageChainTest1(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "2" {
		s.ChannelMessageSend(m.ChannelID, "2")
		s.AddHandlerOnce(messageChainTest2)
	}
}

func messageChainTest2(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "3" {
		s.ChannelMessageSend(m.ChannelID, "3")
	}
}
