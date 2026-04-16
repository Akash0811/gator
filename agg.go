package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"time"
)

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Could not find next feed to fetch")
		return
	}

	dataFeed, err := rss.FetchFeed(
		context.Background(),
		feed.Url,
	)
	if err != nil {
		fmt.Printf("Failed to Fetch from url %s: %v", feed.Url, err)
		return
	}

	fmt.Printf("%-20s | %-40s\n", "Link", "Title")
	fmt.Println("---------------------|------------------------------------------|---------------------")
	fmt.Printf("%-20s | %-40s\n",
		dataFeed.Channel.Link,
		dataFeed.Channel.Title,
		// dataFeed.Channel.Description,
	)

	for _, nestedFeed := range dataFeed.Channel.Item {
		fmt.Printf("%-20s | %-40s\n",
			nestedFeed.Link,
			nestedFeed.Title,
			// nestedFeed.Description,
		)
	}

	err = s.db.MarkFeedFetch(
		context.Background(),
		database.MarkFeedFetchParams{
			UpdatedAt: time.Now(),
			ID:        feed.ID,
		},
	)
	if err != nil {
		fmt.Printf("Could not update feed %s with url %s", feed.Name, feed.Url)
		return
	}

	fmt.Printf("Updated data for feed %s with url %s", feed.Name, feed.Url)
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Agg expects 1 argument")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
