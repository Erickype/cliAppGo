package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("->")
	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	/* 	Remove using len
	   	str = str[0 : len(str)-2] */
	//remove using strings
	str = strings.Replace(str, "\n", "", 1)
	return str, nil
}
