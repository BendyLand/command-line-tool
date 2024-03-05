package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nWelcome to the Go Todo List!")

	var tasks []string
Loop:
	for {
		input := getInput()
		switch input {
		case "add":
			addTask(&tasks)
		case "view":
			viewTasks(tasks)
		case "exit":
			break Loop
		case "delete":
			deleteTask(&tasks)
		case "help":
			fmt.Println("\nThe available commands are:\n\nadd\nview\nexit\ndelete\nhelp\npanic")
			fmt.Println("\nType a command to see more information about it. Or type `back` to exit the help menu.\n")
			var input string
			fmt.Scan(&input)
			showDetails(input)
		case "panic":
			fmt.Println("\nPanicking...")
			os.Exit(1)
		default:
			fmt.Println("\nInvalid command")
		}
	}
	fmt.Println("\nFinal Todo List: ")
	viewTasks(tasks)
}

func showDetails(command string) {
	switch command {
	case "add":
		fmt.Println("\nadd - Adds an item to your list.")
	case "view":
		fmt.Println("\nview - Prints the items currents stored in your list in the order they were added.")
	case "exit":
		fmt.Println("\nexit - Safely stops execution of the program and prints out the final version of your list.")
	case "delete":
		fmt.Println("\ndelete - Prompts the user to select an item to delete by number. A warning is printed if nothing is deleted.")
	case "help":
		fmt.Println("\nhelp - Shows a list of the available commands. Prompts the user for more information.")
	case "panic":
		fmt.Println("\npanic - Stops execution of the program with error code `1`. The final list is NOT displayed like it is with `exit`.")
	case "back":
		fmt.Println("\nHeading back...")
	default:
		fmt.Println("\nInvalid command")
	}
}

func deleteTask(tasks *[]string) {
	var n int
	fmt.Println("\nWhich task would you like to delete?")
	viewTasks(*tasks)
	fmt.Scan(&n)
	n-- // for indexing
	var newTaskList []string
	for i, task := range *tasks {
		if i == n {
			continue
		}
		newTaskList = append(newTaskList, task)
	}
	if len(*tasks) == len(newTaskList) {
		fmt.Println("\nWarning:\nYou may not have entered a valid option.\nPlease double check your list with the `view` command.")
	}
	*tasks = newTaskList
}

func viewTasks(tasks []string) {
	for i, task := range tasks {
		fmt.Printf("\nItem %d.) %s", i+1, task)
	}
	fmt.Println("\n")
}

func addTask(tasks *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nPlease enter the task you would like to add: \n")
	scanner.Scan()
	newTask := scanner.Text()
	newTask = strings.TrimSpace(newTask)
	*tasks = append(*tasks, newTask)
}

func getInput() string {
	var input string
	fmt.Println("\nPlease enter a command or type `help`: \n")
	fmt.Scan(&input)
	return input
}
