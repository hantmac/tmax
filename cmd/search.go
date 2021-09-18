/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/hantmac/tmax/internal/store"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Aliases: []string{"s"},
	Short: "search a command you want",
	Long:  `example: tmax search node, you will get a command that contains 'node'`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("searching and use `Ctrl-c` to exit this program")
		res := GetFuzzySearchResult(strings.Join(args, " "))
		fmt.Println("You may want the following cmd:")

		prompt := promptui.Select{
			Label: "Select your cmd",
			Items: res,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt search exit %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
		execute(result)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func execute(name string) {
	cmd := exec.Command("/bin/sh", "-c", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to execute command: %s\n", err)
		os.Exit(1)
	}
}

func GetFuzzySearchResult(searchStr string) []string {
	s := make([]string, 0)

	st := store.Store()
	for _, v := range st.Shortcuts() {
		s = append(s, v)
	}
	searchResult := fuzzy.Find(searchStr, s)

	return searchResult

}
