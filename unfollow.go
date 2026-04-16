package main

import (
	"context"
	"fmt"

	"github.com/Akash0811/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Unollow expects a single argument")
	}

	feed, err := s.db.GetFeed(
		context.Background(),
		cmd.args[0],
	)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
