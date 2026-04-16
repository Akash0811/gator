package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Following expects no arguments")
	}

	feeds, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		user.Name,
	)
	if err != nil {
		return err
	}

	fmt.Printf("user %s is following the below feeds:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed)
	}

	return nil
}
