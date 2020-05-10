package main

import (
	"log"

	"github.com/chrisaandes/twitter-clone/db"
	"github.com/chrisaandes/twitter-clone/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("mongo no connect")
		return
	}

	handlers.Router()
}
