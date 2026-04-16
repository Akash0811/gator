package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Akash0811/gator/internal/database"
	"github.com/Akash0811/gator/internal/rss"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Could not find next feed to fetch: %v", err)
		return
	}

	dataFeed, err := rss.FetchFeed(
		context.Background(),
		feed.Url,
	)
	if err != nil {
		fmt.Printf("Failed to Fetch from url %s: %v\n", feed.Url, err)
		return
	}

	for _, nestedFeed := range dataFeed.Channel.Item {
		parsedTime, err := time.Parse(time.RFC1123Z, nestedFeed.PubDate)
		if err != nil {
			fmt.Printf("Error while parsing time: %v for %v\n", err, nestedFeed.PubDate)
			continue
		}
		_, err = s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       sql.NullString{String: nestedFeed.Title, Valid: true},
				Url:         nestedFeed.Link,
				Description: sql.NullString{String: nestedFeed.Description, Valid: true},
				PublishedAt: sql.NullTime{Time: parsedTime, Valid: true},
				FeedID:      feed.ID,
			},
		)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
				if pqErr.Code == "23505" {
					fmt.Printf("Already existing post %s with url %s\n", nestedFeed.Title, nestedFeed.Link)
				}
			} else {
				fmt.Printf("Error while posting: %v\n", err)
				return
			}
		}
	}

	err = s.db.MarkFeedFetch(
		context.Background(),
		database.MarkFeedFetchParams{
			UpdatedAt: time.Now(),
			ID:        feed.ID,
		},
	)
	if err != nil {
		fmt.Printf("Could not update feed %s with url %s\n", feed.Name, feed.Url)
		return
	}

	fmt.Printf("Updated data for feed %s with url %s\n", feed.Name, feed.Url)
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
