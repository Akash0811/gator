package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Feeds does not expect arguments")
	}

	rowsFeed, err := s.db.ListFeeds(
		context.Background(),
	)
	if err != nil {
		return err
	}

	fmt.Printf("%-20s | %-40s | %-20s\n", "Name of feed", "URL", "Name of user")
	fmt.Println("---------------------|------------------------------------------|---------------------")

	for _, rowFeed := range rowsFeed {
		fmt.Printf("%-20s | %-40s | %-20s\n",
			rowFeed.Name,
			rowFeed.Url,
			rowFeed.Name_2,
		)
	}

	return nil
}
