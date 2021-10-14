package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		subCommand := os.Args[1]
		if matchSubCommand(subCommand) {
			fmt.Fprintf(os.Stderr, "Error: Use git checkout instead of git %s\n\n", subCommand)
			os.Exit(1)
		}
	}

	if exitCode, err := gitExec(os.Args[1:]); err != nil {
		os.Exit(exitCode)
	}
}

func matchSubCommand(subCommand string) bool  {
	targets := []string{"switch", "restore"}
	for _, target := range targets {
		if target == subCommand {
			return true
		}
	}

	return false
}

func gitExec(args []string) (int, error) {
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		exitCode := cmd.ProcessState.ExitCode()
		return exitCode, err
	}

	return 0, nil
}
