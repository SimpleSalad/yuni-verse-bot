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
	c := i.ApplicationCommandData().Options

	//responses
	rping := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			AllowedMentions: &discordgo.MessageAllowedMentions{
				RepliedUser: true,
			},
			Content: "<@" + i.Member.User.ID + "> Ping Reply!",
		},
	}
	rplain := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Plain Reply",
		},
	}

	//respond with desired response, if none specified default to plain
	if c != nil {
		log.Printf("Test Lamp: %v\n", c[0].IntValue())
		if c[0].IntValue() == 1 {
			s.InteractionRespond(i.Interaction, &rping)
		} else {
			s.InteractionRespond(i.Interaction, &rplain)
		}
	} else {
		log.Println("Test Lamp: No option selected.")
		s.InteractionRespond(i.Interaction, &rplain)
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
