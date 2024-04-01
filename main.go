package main

import (
	database "github.com/abiddiscombe/concierge/internal/database"
	server "github.com/abiddiscombe/concierge/internal/server"
)

func main() {
	database.Init()
	server.Init()
}
