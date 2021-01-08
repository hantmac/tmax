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
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"tmax/internal/core"
	"tmax/setting"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmax",
	Short: "tmax get the cmd you want at lightning speed",
	Long: `The positioning of tmax is a command line tool with a little artificial intelligence. 
If you frequently deal with the terminal daily, tmax will greatly improve your work efficiency.`,
	Args: cobra.ArbitraryArgs,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(core.Args[strings.Join(args, " ")])
			core.Executor(core.Args[strings.Join(args, " ")])
		} else {
			fmt.Println("interactive mode")
			fmt.Printf("tmax %s \n", setting.Version)
			fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
			//interactive mode
			p := prompt.New(
				core.ExecutorForInteractive,
				core.Complete,
				prompt.OptionTitle("tmax: interactive client"),
				prompt.OptionPrefix(">>> "),
				prompt.OptionInputTextColor(prompt.Cyan),
				prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
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

func init() {
	cobra.OnInitialize(initConfig)
	core.Args = core.TransferYamlToMap(setting.FileName)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", setting.FileName, "config file (default is $HOME/.tmax.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tmax.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tmax.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
