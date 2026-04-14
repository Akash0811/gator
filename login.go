package main

import (
	"fmt"
	"gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login expects a single argument")
	}

	err := config.SetUser(*s.cfg, cmd.args[0])
	if err != nil {
		return err
	}
	return nil
}
