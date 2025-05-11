/*
	GShell

	GShell is a minimalist shell implemented in Go, built purely for learning purposes and personal experimentation.
	The goal of this project is to explore how command-line interpreters work under the hood while having fun with Go.
*/

package main

import (
	"bufio"
	"fmt"
	command "gshell/command"
	prompt "gshell/prompt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		prompt.Prompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = command.Exec(input); err != nil {
			fmt.Fprintln(os.Stderr, prompt.Print(err.Error(), prompt.Red))
		}
	}
}
