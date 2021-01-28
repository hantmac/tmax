/*
Copyright © 2020 tmax cloud <hantmac@outlook.com>

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

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/spf13/cobra"

	"tmax/internal/executor"
	"tmax/internal/store"
	"tmax/version"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmax",
	Short: "tmax get the cmd you want at lightning speed",
	Long: `The positioning of tmax is a command line tool with a little artificial intelligence. 
If you frequently deal with the terminal daily, tmax will greatly improve your work efficiency.`,
	Args:               cobra.ArbitraryArgs,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			executor.Execute(args[0], args[1:]...)
		} else {
			fmt.Printf("tmax %s \n", version.Version)
			fmt.Println("Use exit or Ctrl-D (i.e. EOF) to exit")
			//interactive mode
			p := prompt.New(
				func(name string) {executor.Execute(name)},
				complete,
				prompt.OptionTitle("tmax: interactive mode"),
				prompt.OptionPrefix(">>> "),
				prompt.OptionInputTextColor(prompt.Cyan),
				prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
				prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
					if breakline {
						return in == "quit" || in == "exit"
					}
					return false
				}),
			)
			p.Run()
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func complete(d prompt.Document) []prompt.Suggest {

	var s []prompt.Suggest

	t := []prompt.Suggest{
		{Text: "exit", Description: "exit"},
	}

	st := store.Store()

	for k, v := range st.Shortcuts() {
		s = append(s, prompt.Suggest{Text: k, Description: v})
	}

	s = append(s, t...)

	return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
}
