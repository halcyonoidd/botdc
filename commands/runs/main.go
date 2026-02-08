package run 

import (
	"botdc/pkgs/ping"
	"botdc/pkgs/jtc"
	"botdc/pkgs/help"

	"github.com/bwmarrin/discordgo"
)

func RunPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ping.GetPing(s, i)
}