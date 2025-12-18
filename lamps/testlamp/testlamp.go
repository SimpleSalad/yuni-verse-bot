package testlamp

import "github.com/bwmarrin/discordgo"

func commandTest(s *discordgo.Session, i *discordgo.Interaction) {
	s.ChannelMessageSend(i.ChannelID, "Test")
}
