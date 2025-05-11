package command

import (
	"errors"
	"os"
	"strings"
)

func changeDirectory(args []string) error {
	if len(args) < 2 {
		return cdHome()
	}

	return cdCommand(args[1])
}

func cdCommand(dir string) error {
	directory := strings.TrimSpace(dir)

	_, err := os.Stat(directory)
	if err != nil {
		return errors.New("Directory not exist: " + directory)
	}

	return os.Chdir(directory)
}

func cdHome() error {
	if homeDir := os.Getenv("Home"); homeDir == "" {
		return errors.New("Home environment variable not set")
	}

	return os.Chdir("~")
}
