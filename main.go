package main

import (
	database "github.com/abiddiscombe/concierge/internal/database"
	server "github.com/abiddiscombe/concierge/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.Init()
	server.Init()
}
