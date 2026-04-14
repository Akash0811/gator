package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	cfg := config.Read()
	currentState := state{cfg: &cfg}
	currentCommands := commands{cmds: make(map[string]func(*state, command) error)}
	currentCommands.register("login", handlerLogin)

	currentArgs := os.Args
	if len(currentArgs) < 2 {
		fmt.Println("Requires one argument atleast")
		os.Exit(1)
	}

	commandName := currentArgs[1]
	currentArgs = currentArgs[2:]

	err := currentCommands.run(
		&currentState,
		command{name: commandName, args: currentArgs},
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
