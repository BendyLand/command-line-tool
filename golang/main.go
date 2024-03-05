package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Todo List!")
	input := getInput()

	var tasks []string

	switch input {
	case "add":
		addTask(&tasks)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
	for i, item := range tasks {
		fmt.Printf("Item %d.) %s\n", i+1, item)
	}
}

func addTask(tasks *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the task you would like to add: ")
	scanner.Scan()
	newTask := scanner.Text()
	newTask = strings.TrimSpace(newTask)
	*tasks = append(*tasks, newTask)
}

func getInput() string {
	var input string
	fmt.Println("Please enter a command: ")
	fmt.Scan(&input)
	return input
}
