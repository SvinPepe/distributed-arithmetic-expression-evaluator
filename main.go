package main

import (
	"awesomeProject2/parser"

	"bufio"
	"fmt"
	"os"
	"strings"
)

type Expression struct {
	expressionId int
	expression   string
	status       string
}

func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {

	fmt.Print("Enter infix expression: ")
	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error when scanning input:", err.Error())
		return
	}

	fmt.Println("Ya you postfix notation:", parser.ToPostfix(infixString))
	return
}
