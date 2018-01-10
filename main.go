package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Moonlington/discordflo"
	"github.com/bwmarrin/discordgo"
)

var ffs *discordflo.FloFloSession
var conf *Config
var db *sql.DB

func pauseTilDead() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func loadConfig() {
	// Try and load the configuration
	_, err := toml.DecodeFile("config.toml", &conf)
	if os.IsNotExist(err) {
		fmt.Println("No config file found")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//Try to load the config
	loadConfig()
	//And the database
	db = loadDB()

	// Create a new Discord session using the provided bot token
	bot, err := discordflo.New("Bot "+conf.Token, conf.Prefix, false)
	ffs = bot
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		pauseTilDead()
		return
	}

	ffs.AddHandlerOnce(ready)
	ffs.AddHandler(guildMemberChunk)

	// Open a websocket connection to Discord and begin listening
	err = ffs.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		pauseTilDead()
		return
	}

	// Load all commands into discordflo
	loadCommands()
	loadOwnerCommands()

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	pauseTilDead()

	// Cleanly close down the Discord session
	ffs.Close()
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	guilds, _ := s.UserGuilds(100, "", "")
	for _, g := range guilds {
		s.RequestGuildMembers(g.ID, "", 0)
	}
	s.UpdateStatus(0, "| Type "+conf.Prefix+"help |")
}

func guildMemberChunk(s *discordgo.Session, c *discordgo.GuildMembersChunk) {
	for _, g := range s.State.Guilds {
		if g.ID == c.GuildID {
			newm := append(g.Members, c.Members...)
			removeDuplicateMembers(&newm)
			g.Members = newm
			break
		}
	}
}
