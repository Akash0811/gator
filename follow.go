package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Follow expects a single argument")
	}

	feed, err := s.db.GetFeed(
		context.Background(),
		cmd.args[0],
	)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("user %s is now following feed %s!\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
