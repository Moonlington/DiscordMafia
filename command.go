package main

import (
	"fmt"
	"time"

	"github.com/Moonlington/discordflo"
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
			ctx.Sess.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, em)
		}))
}
