package lamps

import (
	"log"

	"github.com/SimpleSalad/yuni-verse-bot/lamps/testlamp"
	"github.com/bwmarrin/discordgo"
)

//add commands

func AddLamps(s *discordgo.Session, g string) {
	testlamp.CreateCommand(s, g)
	log.Println("Adding commands...")
}
