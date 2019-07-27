package Events

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CommandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	if strings.HasPrefix(message.Content, "->b ") {
		textContent := strings.TrimPrefix(message.Content, "->b ")
		textContent = strings.TrimRight(strings.TrimLeft(message.Content, " "), " ")
		textContent = strings.ToLower(textContent)
		if strings.Contains(textContent, "about") || strings.Contains(textContent, "hello") {
			discord.ChannelMessageSend(message.Message.ChannelID, "Hey!,"+message.Message.Author.Username)
		} else {
			discord.ChannelMessageSend(message.Message.ChannelID, "I don't get what you're saying nibba")
		}
	}

	// fmt.Println(content)

	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)

}
