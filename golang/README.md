# Go Todo List

This project is to help me start learning Go. It will be simpler than the other projects in this directory. 

## Usage

1) Run the program with `go run main.go`. You will be prompted to enter a command.
2) While the program is active, you can `add` to, `delete` from, or `view` your todo list; as well as get `help`, safely `exit` the program, or `panic` with an error code. (Prompts are case-sensitive.)
    - *Note: if you enter anything more than the command, then the extra characters will be treated as input to the following prompt.*
    - This can lead to unexpected results if you aren't careful, but it can also be utilized to quickly navigate the interface (e.g. entering "add " will add a blank space as a new task, then immediately return to the prompt asking for a command. However, this behavior also means you can enter things like "add Wash dishes", to select and execute the command all in one step). 

## Motivation

I have tried to learn Go before with small, independent tasks, but I have not tried implementing any standalone projects, even small ones. 

I don't dislike the Go language, but the frustration I have experienced with its toolchain in the past was unrivaled by any other language I've tried, which made it difficult to want to use. This is my attempt to resolve the bitter feeling it previously left. 
