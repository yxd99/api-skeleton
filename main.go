package main

import (
	"api-skeleton/app"
	"log"
)

func main() {

	application := app.NewApp()
	err := application.Run()
	if err != nil {
		log.Fatal("Error to start server")
	}
}
