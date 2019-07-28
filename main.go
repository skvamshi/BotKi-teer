package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/skvamshi/BotKi-teer/config"
	"github.com/skvamshi/BotKi-teer/events"
)

var (
	commandPrefix string
	botID         string
)

func main() {

	// c := scrapper.InitializeScrapper()

	// fmt.Println(c)

	config := config.InitializeConfig()
	discord, err := discordgo.New(config.BotAuthHeader)
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)
	config.BotID = user.ID
	messageService := events.NewTextEvent(config)
	discord.AddHandler(messageService.CommandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "A friendly helpful bot!")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("Bot Ki'teer has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	// commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}
