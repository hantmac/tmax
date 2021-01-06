package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"tmax/internal/debug"
)

// ExecutorForInteractive find real cmd and exec it
func ExecutorForInteractive(s string) {

	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	Executor(Args[s])
}

func Executor(s string) {

	fmt.Println(s)

	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
