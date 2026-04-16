package main

import (
	"context"
	"fmt"
	"gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Agg does not expect arguments")
	}

	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(
		context.Background(),
		url,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Data: %v\n\n", *feed)
	return nil
}
