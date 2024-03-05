package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Todo List!")
    
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
            fmt.Println("Panicking...")
            os.Exit(1)
        default:
            fmt.Println("Invalid command")
        }
    }
	viewTasks(tasks)
}

func deleteTask(tasks *[]string) {
    var n int
    fmt.Println("Which task would you like to delete?")
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
    *tasks = newTaskList
}

func viewTasks(tasks []string) {
    for i, task := range tasks {
        fmt.Printf("Item %d.) %s\n", i+1, task)
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
