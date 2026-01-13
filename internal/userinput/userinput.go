package userinput

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// TODO: in learning purposes do with two different variants, with user prompting and with cli flags

// Read user input and return as string
func GetUserInput() string {

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input

}
