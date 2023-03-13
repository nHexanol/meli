package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// data passed from exec command
	// if length isnt 8 then something is probably wrong
	if len(os.Args) != 8 {
		panic("Error in `embed.go` at func `main:17:6`")
	}

	var userId = os.Args[1]
	var nameAge = os.Args[2]
	var edu = os.Args[3]
	var wars = os.Args[4]
	var eco = os.Args[5]
	var fav = os.Args[6]
	var thumbnailsUrl = os.Args[7]

	if len(fav) == 0 {
		fav = "<nil>"
	}

	var t = loadToken()
	var client, e = discordgo.New("Bot " + t)
	if e != nil {
		panic(e)
	}

	var embed = &discordgo.MessageEmbed{
		Title:       "Guild application",
		Color:       0xFF8888,
		Description: fmt.Sprintf("From <@%s>", userId),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Username and age",
				Value:  nameAge,
				Inline: false,
			},
			{
				Name:   "Education and previous guild affiliation",
				Value:  edu,
				Inline: false,
			},
			{
				Name:   "Guild wars and eco",
				Value:  wars,
				Inline: false,
			},
			{
				Name:   "Weekly activities and classes",
				Value:  eco,
				Inline: false,
			},
			{
				Name:   "Favourite things to do on Wynncraft",
				Value:  fav,
				Inline: false,
			},
		},
		Timestamp: "",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: thumbnailsUrl,
		},
	}

	client.ChannelMessageSendEmbed("1073847845553176596", embed)

	client.Close()

	os.Exit(0)
}

func loadToken() string {
	var token, err = ioutil.ReadFile("./token.txt")
	if err != nil {
		panic(err)
	}
	token = bytes.TrimSpace(token)
	var tokenString = string(token)
	return tokenString
}
