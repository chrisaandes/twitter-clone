package main

import (
	"log"

	"github.com/chrisaandes/twitter-clone/db"
	"github.com/chrisaandes/twitter-clone/handlers"
	"github.com/chrisaandes/twitter-clone/handlers/handlers.go"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("mongo no connect")
		return
	}

	handlers.Router()
}
