package main

import (
	"log"

	"diary.mingky.me/database"
	"diary.mingky.me/route"
	"diary.mingky.me/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.Create()
	database.SetupDatabase()
	route.SetupRoute(app)
	app.Listen(":80")

}
