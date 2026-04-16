package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Akash0811/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("Unollow expects atmost single argument")
	}
	numPosts := 2
	if len(cmd.args) == 1 {
		num, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Browse takes a integer as argument (optionally)")
		}
		numPosts = num
	}

	posts, err := s.db.GetPost(
		context.Background(),
		database.GetPostParams{
			Limit:  int32(numPosts),
			UserID: user.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("%-20s | %-40s | %-40s\n", "Name of Feed", "Link", "Title")
	fmt.Println("---------------------|------------------------------------------|---------------------")
	for _, post := range posts {
		fmt.Printf("%-20s | %-40s | %-40s\n",
			post.FeedName,
			post.Url,
			post.Title.String,
		)
	}

	return nil
}
