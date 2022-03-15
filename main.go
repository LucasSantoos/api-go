package main

import (
	"github.com/LucasSantoos/api-go/database"
	"github.com/LucasSantoos/api-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
