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
	"github.com/google/martian/log"
	"os"
	"strings"
	"tmax/internal/core"
	"tmax/setting"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command to tmax.yaml",
	Long: `add a command to tmax.yaml, the format is "group:cmdKey:cmdValue", 
example: k8s:getpod:kubectl get pod
`,
// support add cmd like k8s:getpod:kubectl get pod
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		s := strings.Join(args, " ")
		fmt.Println(s)
		ss := strings.Split(s, ":")
		if len(ss) == 3 {
			m := make(map[string]string)
			m[ss[1]] = ss[2]
			outMap := core.TransYamlToOutMap(setting.FileName)
			outMap.InsertToMap(ss[0], m)

			b, err := core.TransMapToYaml(outMap)
			if err != nil {
				fmt.Printf("transfer map to yaml failed, %v", err)
				return
			}

			// re-write tmax.yaml
			f, err := os.OpenFile(setting.FileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)

			if err != nil {
				log.Errorf("open .tmax.yaml failed, %v", err)
				return
			}
			_, err = f.Write(b)
			if err != nil {
				log.Errorf("write .tmax.yaml failed, %v", err)
			}

		} else {
			fmt.Println("add cmd failed")
		}

		fmt.Printf("cmd %s add to tmax.yaml success", strings.Join(ss[1:], " "))

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
