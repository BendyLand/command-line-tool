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
    Loop: for {
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
        case "panic":
            fmt.Println("\nPanicking...")
            os.Exit(1)
        default:
            fmt.Println("\nInvalid command")
        }
    }
    fmt.Println("Final Todo List:")
	viewTasks(tasks)
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
	fmt.Println("\nPlease enter a command: \n")
	fmt.Scan(&input)
	return input
}
