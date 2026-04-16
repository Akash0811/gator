package main

import (
	"context"
	"fmt"

	"github.com/Akash0811/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login expects a single argument")
	}

	_, err := s.db.GetUser(
		context.Background(),
		cmd.args[0],
	)
	if err != nil {
		return err
	}

	err = config.SetUser(*s.cfg, cmd.args[0])
	if err != nil {
		return err
	}
	return nil
}
