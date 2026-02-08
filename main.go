package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		panic("Error creating Discord session: " + err.Error())
	}

	if err := dg.Open(); err != nil {
		panic("Error opening connection: " + err.Error())
	}
	defer dg.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}




func RegisterCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with a ping!",
		},
		{
			Name:        "help",
			Description: "Get help with commands.",
		},
	}
	
	appID := s.State.User.ID // This is your application's ID

	// Loop through and register each command
	for _, cmd := range commands {
		// Register the command globally (empty string for GuildID means global)
		if _, err := s.ApplicationCommandCreate(appID, "", cmd); err != nil {
			fmt.Println("Cannot create command:", cmd.Name, err)
		}
	}
}