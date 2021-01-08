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
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strings"
	"tmax/internal/core"
)

// sCmd represents the s command
var sCmd = &cobra.Command{
	Use:   "s",
	Short: "search a command you want",
	Long:  `example: tmax s node, you will get a command that contains 'node'`,
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
		core.Executor(result)
	},
}

func init() {
	rootCmd.AddCommand(sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
