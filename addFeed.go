package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("AddFeed expects two arguments. Name of Feed and URL in that order...")
	}

	currentUser, err := s.db.GetUser(
		context.Background(),
		s.cfg.CurrentUserName,
	)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
			Url:       cmd.args[1],
			UserID:    currentUser.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Feed %v was registered successfully!\n", feed)

	return nil
}
