package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Following expects no arguments")
	}

	feeds, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		s.cfg.CurrentUserName,
	)
	if err != nil {
		return err
	}

	fmt.Printf("user %s is following the below feeds:\n", s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed)
	}

	return nil
}
