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

	"github.com/mirrorhub-io/platform/client"
	"github.com/spf13/cobra"
	"log"
)

var (
	name     string
	email    string
	password string
)

var clientContactCmd = &cobra.Command{
	Use:   "contact",
	Short: "mirrorhub contacts utils",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
	},
}
var clientContactCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "mirrorhub contact create",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.Initialize()
		x, err := c.Contact().Create(name, email, password)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(x)
		fmt.Println("client called")
	},
}

func init() {
	RootCmd.AddCommand(clientContactCmd)
	clientContactCmd.AddCommand(clientContactCreateCmd)
	clientContactCreateCmd.Flags().StringVarP(&name, "name", "n", "", "contact name")
	clientContactCreateCmd.Flags().StringVarP(&email, "email", "e", "", "contact email")
	clientContactCreateCmd.Flags().StringVarP(&password, "password", "p", "", "contact password (min 8 chars)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
