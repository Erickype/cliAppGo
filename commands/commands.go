package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Erickype/cliAppGo/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("-> ")
	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	/*Remove using len
	str = str[0 : len(str)-2] */
	//remove using strings
	str = strings.Replace(str, "\r\n", "", 1)
	/* for i := 0; i < len(str); i++ {
		fmt.Printf("Index: %d Value:%d \n", i, str[i])
	} */
	return str, nil
}

func ShowInConsole(expensesList []float32) {
	var builder = strings.Builder{}

	fmt.Println("List of expenses")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		if i == len(expensesList)-1 {
			fmt.Println("====================")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", expenses.Sum(expensesList...)))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", expenses.Max(expensesList...)))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", expenses.Min(expensesList...)))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", expenses.Average(expensesList...)))
		}
	}

	fmt.Println(builder.String())
}

func Export() {

}
