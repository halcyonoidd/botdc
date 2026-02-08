package run

import (
	"botdc/pkgs/help"
	"botdc/pkgs/jtc"
	"botdc/pkgs/ping"

	"github.com/bwmarrin/discordgo"
)

func RunPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ping.GetPing(s, i)
}

func RunJTC(s *discordgo.Session, i *discordgo.InteractionCreate) {
	jtc.GetJTC(s, i)
}

func RunHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	help.GetHelp(s, i)
}