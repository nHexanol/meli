package main

import (
	"os"
	"time"
)

func CacheTerritoryAsync() {
	go func() {
		for {
			// do shit later xd
			os.Rename("./cache/territory.json", "./cache/territory.json.bak")
			time.Sleep(24 * time.Hour)
		}
	}()
}
