package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ApiPort int
}

var Env Config

func StartConfig() error {
	apiPortStr := os.Getenv("API_PORT")
	if apiPortStr == "" {
		return fmt.Errorf("API_PORT not defined")
	}

	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		return fmt.Errorf("failed to convert API_PORT to int: %w", err)
	}

	Env = Config{
		ApiPort: apiPort,
	}

	return nil
}
