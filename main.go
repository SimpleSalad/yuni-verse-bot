package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func init() {
	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	
	dg, err := discordgo.New("Bot " + BotToken)
	
}