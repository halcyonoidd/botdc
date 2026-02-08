package main

import (
	"fmt"
	"os"

	"botdc/commands/runs"

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

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.ApplicationCommandData().Name {
		case "ping":
			runs.RunPing(s, i) // Memanggil fungsi dari commands/runs/main.go
		case "help":
			runs.RunHelp(s, i) // Memanggil fungsi dari commands/runs/main.go
		// case "jtc":           // Jika Anda ingin mengaktifkan JTC juga
		// 	runs.RunJTC(s, i)
		}
	})

	if err := dg.Open(); err != nil {
		panic("Error opening connection: " + err.Error())
	}
	defer dg.Close()

	RegisterCommands(dg)

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
		{ // Jangan lupa daftarkan command jtc jika ingin dipakai
			Name:        "jtc",
			Description: "Join to create channel command",
		},
	}

	appID := s.State.User.ID

	for _, cmd := range commands {
		if _, err := s.ApplicationCommandCreate(appID, "", cmd); err != nil {
			fmt.Println("Cannot create command:", cmd.Name, err)
		} else {
			fmt.Println("Command created:", cmd.Name)
		}
	}
}