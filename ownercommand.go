package main

import (
	"math/rand"

	"github.com/Moonlington/discordflo"
)

func loadOwnerCommands() {

	// ffs.AddPrivateCommand("Owner", func(ctx *discordflo.Context) bool {
	// 	if ctx.Mess.Author.ID == conf.Owner {
	// 		return true
	// 	}
	// 	return false
	// }, discordflo.NewCommand(
	// 	"Name",
	// 	"Desc",
	// 	"Usage",
	// 	"Detailed",
	// 	func(ctx *discordflo.Context) {
	//
	// 	}))

	ffs.AddPrivateCommand("Game Control", func(ctx *discordflo.Context) bool {
		if ctx.Mess.Author.ID == conf.Owner {
			return true
		}
		return false
	}, discordflo.NewCommand(
		"gencities",
		"",
		"",
		"",
		func(ctx *discordflo.Context) {
			for i := 1; i <= 10; i++ {
				addCity(cityName())
			}
		}))

	ffs.AddPrivateCommand("Game Control", func(ctx *discordflo.Context) bool {
		if ctx.Mess.Author.ID == conf.Owner {
			return true
		}
		return false
	}, discordflo.NewCommand(
		"genlocations",
		"",
		"",
		"",
		func(ctx *discordflo.Context) {
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
		}))
}
