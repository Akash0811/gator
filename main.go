package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"os"

	_ "github.com/lib/pq" // Postgres driver
)

func main() {
	cfg := config.Read()

	db, err := sql.Open("postgres", cfg.DBUrl)
	dbQueries := database.New(db)

	currentState := state{cfg: &cfg, db: dbQueries}

	currentCommands := commands{cmds: make(map[string]func(*state, command) error)}
	currentCommands.register("login", handlerLogin)
	currentCommands.register("register", handlerRegister)
	currentCommands.register("reset", handlerReset)
	currentCommands.register("users", handlerListUsers)
	currentCommands.register("agg", handlerAgg)
	currentCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	currentCommands.register("feeds", handlerListFeeds)
	currentCommands.register("follow", middlewareLoggedIn(handlerFollow))
	currentCommands.register("following", middlewareLoggedIn(handlerFollowing))
	currentCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	currentArgs := os.Args
	if len(currentArgs) < 2 {
		fmt.Println("Requires one argument atleast")
		os.Exit(1)
	}

	commandName := currentArgs[1]
	currentArgs = currentArgs[2:]

	err = currentCommands.run(
		&currentState,
		command{name: commandName, args: currentArgs},
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
