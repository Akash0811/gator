package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Users does not expect arguments")
	}

	nameUsers, err := s.db.GetUsers(
		context.Background(),
	)
	if err != nil {
		return err
	}

	currentNameUser := s.cfg.CurrentUserName
	for _, nameUser := range nameUsers {
		if nameUser == currentNameUser {
			fmt.Printf("* %v (current)\n", nameUser)
		} else {
			fmt.Printf("* %v\n", nameUser)
		}
	}

	return nil
}
