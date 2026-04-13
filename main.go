package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg := config.Read()
	if cfg == (config.Config{}) {
		return
	}
	err := config.SetUser(cfg, "cute")
	if err != nil {
		return
	}
	newCfg := config.Read()
	if newCfg == (config.Config{}) {
		return
	}
	fmt.Printf("Current Configuration: %v\n", newCfg)
}
