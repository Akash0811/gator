package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Read() Config {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("%v\n", err)
		return Config{}
	}
	file, err := os.Open(cfgPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return Config{}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return Config{}
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("%v\n", err)
		return Config{}
	}

	return cfg
}

func SetUser(cfg Config, CurrentUserName string) error {
	cfg.CurrentUserName = CurrentUserName
	err := write(cfg)
	if err != nil {
		return fmt.Errorf("Unable to write to file!")
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to locate Home Directory")
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(cfgPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
