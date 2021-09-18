package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hantmac/tmax/internal/store"
)

func Execute(name string, args ...string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	if name == "quit" || name == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}

	s := store.Store()
	longName, ok := s.GetFullCommand(name)
	if !ok {
		fmt.Printf("command not found: %s\n", name)
		os.Exit(1)
	}

	name = longName

	if len(args) > 0 {
		if hasTemplate(name) {
			argsMap, err := buildArgs(args)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to build args: %s\n", err)
				os.Exit(1)
			}

			name, err = parseCommand(name, argsMap)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to parse command: %s\n", err)
				os.Exit(1)
			}
		} else {
			name = fmt.Sprintf("%s %s", name, strings.Join(args, " "))
		}
	}

	fmt.Println(name)

	cmd := exec.Command("/bin/sh", "-c", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to execute command: %s\n", err)
		os.Exit(1)
	}
}
