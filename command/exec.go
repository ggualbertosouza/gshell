package command

import (
	"os"
	"os/exec"
	"strings"
)

func Exec(input string) error {
	input = strings.TrimSuffix(input, "\n")
	pipeline := strings.Split(input, "|")

	if len(pipeline) > 1 {
		for _, command := range pipeline {
			execCommand(command)
		}
	}

	return execCommand(input)
}

func execCommand(input string) error {
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		return changeDirectory(args)
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
