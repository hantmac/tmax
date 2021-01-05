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
	"github.com/google/martian/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"path"
	"tmax/internal/core"
)

var tmaxDefaultConf = `
---
k8s:
  filternodecpu: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.capacity.cpu}{'\t'}{.status.capacity.memory}{'\n'}{end}"
  filternodetaint: kubectl get nodes -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.spec.taints[*].key}{'\n'}{end}"
  corednsedit: kubectl edit cm coredns -nkube-system
  allnode: kubectl get no
  alldeploy: kubectl get deploy
  allpod: kubectl get pod -A
  busyboxrun: kubectl run busybox --rm -ti --image=busybox /bin/sh
  allnodeip: kubectl get node -o=jsonpath="{range .items[*]}{.metadata.name}{'\t'}{.status.addresses[0].address}{'\n'}{end}"
  podResource: kubectl get pod -o custom-columns=NAME:metadata.name,podIP:status.podIP,hostIp:spec.containers[0].resources | grep 8Gi
custom:
  check: curl 127.0.0.1:8080/@in/api/health
unix:
  '': ''
`

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate default cmd config",
	Long:  `generate default cmd config`,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateTmaxYaml()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GenerateTmaxYaml() {
	homedirStr, err := homedir.Dir()
	if err != nil {
		log.Errorf("get home dir failed: %v", err)
	}
	fileName := path.Join(homedirStr, ".tmax.yaml")
	if core.ExistFile(fileName) {
		log.Errorf("the .tmax.yaml already exist")
		return
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Errorf("open tmax.yaml failed")
		return
	} else {
		_, err = f.Write([]byte(tmaxDefaultConf))
	}
	defer f.Close()

	fmt.Println(" tmax.yaml generated in ~/.tmax.yaml")
}
