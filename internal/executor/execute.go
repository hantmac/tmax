package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"

	"tmax/internal/store"
)

func Execute(name string, args ...string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	} else if name == "quit" || name == "exit" {
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
		argsMap, err := buildArgs(args)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to build args: %s\n", err)
			os.Exit(1)
		}

		origin := name
		name, err = parseCommand(name, argsMap)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to parse command: %s\n", err)
			os.Exit(1)
		}

		if name == origin {
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

// buildArgs accepts a series of arguments and generates a key-value map.
// e.g. "-a b -c=d -e --f g" => {"a": "b", "c": "d", "e": "true", "f": "g"}
func buildArgs(args []string) (map[string]string, error) {
	fieldsCount := len(args)
	res := make(map[string]string, fieldsCount)
	for i := 0; i < fieldsCount; {
		arg := args[i]
		if !isNewFlag(arg) {
			return nil, fmt.Errorf("wrong params")
		}

		k, v, ok := extractKV(arg)
		if ok {
			res[k] = v
			i += 1
			continue
		}

		if i == fieldsCount-1 {
			res[trimArg(arg)] = "true"
			break
		}

		nextArg := args[i+1]
		if isNewFlag(nextArg) {
			res[trimArg(arg)] = "true"
			i += 1
			continue
		}

		res[trimArg(args[i])] = nextArg
		i += 2
	}

	return res, nil
}

func isNewFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

func extractKV(arg string) (string, string, bool) {
	kv := strings.Split(arg, "=")
	if len(kv) != 2 {
		return arg, "", false
	}

	return trimArg(kv[0]), kv[1], true
}

// trimArg removes the "-" or "--" prefix of the given argument.
func trimArg(arg string) string {
	return strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-")
}

func parseCommand(name string, args map[string]string) (string, error) {
	var b strings.Builder

	tmpl, err := template.New("tmpl").Funcs(template.FuncMap(sprig.FuncMap())).Parse(name)

	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&b, args)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
