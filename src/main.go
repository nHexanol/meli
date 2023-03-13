package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	guild "meli/helper"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	var t = loadToken()
	var client, e = discordgo.New("Bot " + t)
	if e != nil {
		panic(e)
	}

	if ee := recover(); ee != nil {
		client.ChannelMessageSend("1058954802018648114", "```golang\n"+e.Error()+"\n```")
	}

	sh := make(chan os.Signal, 1)
	signal.Notify(sh, syscall.SIGSEGV, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	go func() {
		select {
		case <-sh:
			var errmsg = "```golang\nInterrupted\n\n" + string(debug.Stack())
			client.ChannelMessageSend("1058954802018648114", errmsg+"\n```")
			fmt.Println(string(debug.Stack()))
		}
	}()

	client.AddHandler(message)

	client.Identify.Intents = discordgo.IntentsGuildMessages

	err2 := client.Open()
	if err2 != nil {
		panic(err2)
	}
	getTerritoryList(client)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

}

// message handler
func message(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.Author.Username, m.Content)
}

// token loader from file
func loadToken() string {
	var token, err = ioutil.ReadFile("./token.txt")
	if err != nil {
		panic(err)
	}
	token = bytes.TrimSpace(token)
	var tokenString = string(token)
	return tokenString
}

func getTerritoryList(client *discordgo.Session) {
	var resp, err = http.Get("https://api.wynncraft.com/public_api.php?action=territoryList")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var territories map[string]map[string]*guild.Territory
	err = json.Unmarshal(body, &territories)
	if err != nil {
		panic(err)
	}

	fmt.Println(territories["territories"]["Ragni"])
	return
}
