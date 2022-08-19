package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/Erickype/cliAppGo/commands"
)

func main() {

	var expenses []float32
	var export string
	flag.StringVar(&export, "export", "", "Exports details of expenses")
	flag.Parse()

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
			continue
		}

		expenses = append(expenses, float32(expense))
	}

	if export == "" {
		commands.ShowInConsole(expenses)

	} else {
		commands.Export(export, expenses)
	}

}
