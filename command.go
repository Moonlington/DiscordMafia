package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Moonlington/discordflo"
	"github.com/bwmarrin/discordgo"
)

func loadCommands() {

	// ffs.AddCommand("Category", discordflo.NewCommand(
	// 	"Name",
	// 	"Desc",
	// 	"Usage",
	// 	"Detailed",
	// 	func(ctx *discordflo.Context) {
	//
	// 	}))

	ffs.AddCommand("Bot", discordflo.NewCommand(
		"ping",
		"Pings the bot and returns the response time",
		"[ex]",
		"The optional argument `ex` tests extra pinging",
		func(ctx *discordflo.Context) {
			em := createEmbed(ctx, "Ping!")
			now := time.Now()
			msg, err := ctx.SendEmbed(em)
			if err != nil {
				return
			}
			then := time.Since(now)
			em.Description = fmt.Sprintf("Ping! `%s`", then.String())
			if ctx.Argstr == "ex" {
				now2 := time.Now()
				ffs.MessageReactionAdd(msg.ChannelID, msg.ID, "💥")
				then2 := time.Since(now2)
				em.Description += fmt.Sprintf(" | Reaction: `%s`", then2.String())
				now3 := time.Now()
				addReactionQueue(func(m *discordgo.MessageReactionAdd) bool {
					if m.MessageID == msg.ID && m.Emoji.Name == "💥" && m.UserID != ffs.State.User.ID {
						then3 := time.Since(now3)
						em.Description += fmt.Sprintf(" | You: `%s`", then3.String())
						ffs.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, em)
						return true
					}
					return false
				})
			}
			ffs.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, em)
		}))

	ffs.AddCommand("Owner", discordflo.NewCommand(
		"gencities",
		"",
		"",
		"",
		func(ctx *discordflo.Context) {
			if ctx.Mess.Author.ID == conf.Owner {
				for i := 1; i <= 10; i++ {
					addCity(cityName())
				}
			}
		}))
	ffs.AddCommand("Owner", discordflo.NewCommand(
		"genlocations",
		"",
		"",
		"",
		func(ctx *discordflo.Context) {
			if ctx.Mess.Author.ID == conf.Owner {
				var amount int
				err := db.QueryRow("SELECT COUNT(ID) FROM Cities").Scan(&amount)
				if err != nil {
					return
				}
				for i := 1; i <= amount; i++ {
					for j := 1; j <= 25; j++ {
						d := rand.Float64() * 10
						w := 1575 * d
						addCompany(companyName(), w, d, i)
					}
				}
			}
		}))
}
