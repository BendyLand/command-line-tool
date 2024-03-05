package main

import (
    "fmt"
    "os"
)

func main() {
	fmt.Println("Welcome to the Go Todo List!")
	input := getInput()

	switch input {
		case "add": 
			fmt.Println("You typed add")
		default: 
			fmt.Println("Invalid command")
			os.Exit(1)
	}
}

func getInput() string {
	var input string
	fmt.Println("Please enter a command: ")
	fmt.Scan(&input)
	return input
}


