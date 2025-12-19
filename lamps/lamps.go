package lamps

import (
	"github.com/SimpleSalad/yuni-verse-bot/lamps/testlamp"
	"github.com/bwmarrin/discordgo"
)

//add commands

func AddLamps(s *discordgo.Session, g string) {
	testlamp.CreateCommand(s, g)
}
