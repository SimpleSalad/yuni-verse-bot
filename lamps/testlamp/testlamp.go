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
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "type",
				Description: "The type of response to request from the bot. Default: plain.",
				Required:    true,
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
	c := i.ApplicationCommandData().Options
	log.Println(c[0].IntValue())
	if c[0].IntValue() == 1 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				AllowedMentions: &discordgo.MessageAllowedMentions{
					RepliedUser: true,
				},
				Content: "<@" + i.Member.User.ID + "> Ping Reply!",
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Plain Reply",
			},
		})
	}
}

func CreateCommand(s *discordgo.Session, g string) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, g, &cmd)
	log.Printf("Creating command '%v'...\n", cmd.Name)
	if err != nil {
		log.Panicf("Error creating command '%v': %v\n", &cmd.Name, err)
	}

	s.AddHandler(commandTest)
}
