package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func ShowInConsole(expenses []float32) {
	var builder = strings.Builder{}

	for _, expense := range expenses {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
	}
}
