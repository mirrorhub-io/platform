// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"errors"
	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	autocompleteTarget string
	autocompleteType   string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mirrorhub",
	Short: "Mirrorhub root command.",
}

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate shell autocompletion script for Mirrorhub",
	RunE: func(cmd *cobra.Command, args []string) error {
		if autocompleteType != "bash" {
			return errors.New("Only Bash is supported for now")
		}
		err := cmd.Root().GenBashCompletionFile(autocompleteTarget)
		if err != nil {
			return err
		}
		color.Green("Bash completion file for Mirrorhub saved to: " + autocompleteTarget)
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mirrorhub.yaml)")
	RootCmd.AddCommand(autocompleteCmd)

	autocompleteCmd.PersistentFlags().StringVarP(&autocompleteTarget, "completionfile", "", "/etc/bash_completion.d/mirrorhub.sh", "Autocompletion file")
	autocompleteCmd.PersistentFlags().StringVarP(&autocompleteType, "type", "", "bash", "Autocompletion type (currently only bash supported)")
	autocompleteCmd.PersistentFlags().SetAnnotation("completionfile", cobra.BashCompFilenameExt, []string{})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".mirrorhub") // name of config file (without extension)
	viper.SetEnvPrefix("MIRRORHUB")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".config")
	viper.AddConfigPath("/etc/mirrorhub")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.ReadInConfig()

	viper.SetDefault("Email", "admin@mirrorhub.io")
	viper.SetDefault("Password", "")
	viper.SetDefault("API", map[string]string{
		"gateway": "localhost:8080",
		"base":    "localhost:9000",
	})

}

func exit(msg string) {
	color.Red(msg)
	os.Exit(1)
}

func ret(val interface{}, err error) {
	s, _ := prettyjson.Marshal(val)
	if err != nil {
		s, _ = prettyjson.Marshal(err.Error())
		color.Red(string(s))
		os.Exit(1)
	}
	fmt.Println(string(s))
	os.Exit(0)
}
