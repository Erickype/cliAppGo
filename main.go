package main

import (
	"log"
	"strconv"

	"github.com/Erickype/cliAppGo/commands"
)

func main() {

	var expenses []float32

	for {
		input, err := commands.GetInput()

		if err != nil {
			log.Fatalln(err)
		}

		if input == "n" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			log.Fatalln(err)
		}

		expenses = append(expenses, float32(expense))
	}

	commands.ShowInConsole(expenses)
}
