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
