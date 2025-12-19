package lamps

import (
	"github.com/bwmarrin/discordgo"
)

//add commands

func AddLamps(s *discordgo.Session, g string) {
	var cmd = *testlamp.Command
	s.ApplicationCommandCreate(s.State.User.ID, g, cmd)
}
