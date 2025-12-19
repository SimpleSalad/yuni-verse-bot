package lamps

import (
	"github.com/bwmarrin/discordgo"
)

//add commands

func addLamps(s *discordgo.Session, g string) {
	var cmd = *testlamp.cmd
	s.ApplicationCommandCreate(s.State.User.ID, g, cmd)
}
