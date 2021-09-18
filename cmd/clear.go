/*
Copyright © 2021 tmax cloud <hantmac@outlook.com>

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

	"github.com/hantmac/tmax/setting"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove .tmax.yaml",
	Long:  `remove .tmax.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteConfig()
		fmt.Println("remove .tmax.yaml succeeded, you can re-generate it by running 'tmax generate'")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func deleteConfig() {
	err := os.Remove(setting.ConfigPath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("can not delete file %s: %s\n", setting.ConfigPath, err)
		os.Exit(1)
	}
}
