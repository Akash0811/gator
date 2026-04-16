package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login expects a single argument")
	}

	_, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
		},
	)
	if err != nil {
		return err
	}

	err = config.SetUser(*s.cfg, cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("user %s was created successfully!\n", cmd.args[0])

	return nil
}
