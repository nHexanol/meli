package main

import (
	"fmt"
	"os"
)

func main() {
	// data passed from argv
	// first args is username, second is discord id
	var uuid = os.Args[1]
	var discordId = os.Args[2]

	var data = []byte(fmt.Sprintf("{\"%s\": \"%s\"},\n", uuid, discordId))
	fmt.Printf("Linking %s to %s", uuid, discordId)
	// append to file not overwriting

	f, err := os.OpenFile("../data/userData.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := f.Write(data); err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	os.Exit(0)
}
