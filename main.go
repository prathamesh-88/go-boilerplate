package main

import (
	"log"

	"github.com/prathamesh-88/go-boilerplate/pkg/database"
	"github.com/prathamesh-88/go-boilerplate/routes"
)

func main() {
	database.Connect()
	defer database.CloseConnection()

	server := routes.Server{}
	server.Port = 3000

	err := server.Start()

	if err != nil {
		log.Fatal("Failed to start the server ", err.Error())
	}
}
