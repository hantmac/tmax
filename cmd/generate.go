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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/hantmac/tmax/setting"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate default config",
	Long:  `generate default config, it will override the current config if existing`,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateConfig()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func GenerateConfig() {
	err := ioutil.WriteFile(setting.ConfigPath, []byte(setting.DefaultConf), os.ModePerm)
	if err != nil {
		fmt.Printf("can not generate file %s\n", setting.ConfigPath)
		os.Exit(1)
	}
	fmt.Println(".tmax.yaml generated in ~/.tmax.yaml")
}
