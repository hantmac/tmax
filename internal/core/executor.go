package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
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

func Executor(s string, args ...string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	if len(args) > 0 {
		argsMap, err := buildArgs(args)
		if err != nil {
			fmt.Printf("Failed to build args: %s\n", err)
		}

		old := s
		s, err = parseCommand(s, argsMap)

		if s == old {
			s = fmt.Sprintf("%s %s", s, strings.Join(args, " "))
		}
	}

	fmt.Println(s)

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

func buildArgs(args []string) (map[string]string, error) {
	fieldsCount := len(args)
	if fieldsCount/2 == 0 {
		return nil, fmt.Errorf("wrong params")
	}

	res := make(map[string]string)
	for i := 0; i < fieldsCount; i += 2 {
		res[trimArgs(args[i])] = args[i+1]
	}

	return res, nil
}

func trimArgs(s string) string {
	return strings.TrimPrefix(strings.TrimPrefix(s, "-"), "-")
}

func parseCommand(input string, args map[string]string) (string, error) {
	var b strings.Builder

	tmpl, err := template.New("tmpl").Parse(input)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&b, args)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
