package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Reset does not expect arguments")
	}

	err := s.db.TruncateUsers(
		context.Background(),
	)
	if err != nil {
		return err
	}

	fmt.Println("Users table emptied successfully!")

	return nil
}
