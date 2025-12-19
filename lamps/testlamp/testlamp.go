package testlamp

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	cmd = discordgo.ApplicationCommand{
		Name:        "test",
		Description: "sends a test reply",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        3,
				Name:        "type",
				Description: "The type of response to request from the bot. Default: plain.",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Plain",
						Value: 0,
					},
					{
						Name:  "Ping",
						Value: 1,
					},
				},
			},
		},
	}
)

// send a message in chat
func commandTest(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.ChannelMessageSend(i.ChannelID, "Test")
}

func CreateCommand(s *discordgo.Session, g string) {
	ccmd, err := s.ApplicationCommandCreate(s.State.User.ID, g, &cmd)
	log.Printf("Creating command '%v'...\n", &ccmd.Name)
	if err != nil {
		log.Panicf("Error creating command '%v': %v\n", &cmd.Name, err)
	}

	s.AddHandler(commandTest)
}
