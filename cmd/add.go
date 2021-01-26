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

	"github.com/spf13/cobra"

	"tmax/internal/store"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <short command> <full command>",
	Short: "add command to tmax.yaml",
	Long: "add a command to tmax.yaml. <short command> is grouped, a command named 'a.b' means adding command b to group a",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		s := store.Store()
		err := s.AddCommand(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("command added")

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
