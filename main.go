package main

import (
	"log"

	"github.com/Erickype/cliAppGo/commands"
)

func main() {

	for {
		input, err := commands.GetInput()

		if err != nil {
			log.Fatalln(err)
		}

		if input == "n" {
			break
		}
	}
	log.Println("Finalizado")
}
