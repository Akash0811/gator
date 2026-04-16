package main

import (
	"github.com/Akash0811/gator/internal/config"
	"github.com/Akash0811/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
