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
	"github.com/mirrorhub-io/platform/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os/user"
)

var (
	name     string
	email    string
	password string
	token    string
	c        *client.Client
)

func currentTokenFile() string {
	u, _ := user.Current()
	return "/tmp/.mirrorhub." + u.Username + ".token"
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Mirrorhub API-Client",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		c = client.Initialize()
		r, _ := ioutil.ReadFile(currentTokenFile())
		if len(r) > 0 {
			c.ContactToken = string(r)
			c.PrepareHeader()
		}
		if email == "" {
			email = viper.GetString("Email")
		}
		if password == "" {
			password = viper.GetString("Password")
		}
	},
}

var clientContactCmd = &cobra.Command{
	Use:   "contact",
	Short: "mirrorhub contacts utils",
}

var clientContactCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create contact",
	Run: func(cmd *cobra.Command, args []string) {
		co, err := c.Contact().Create(name, email, password)
		if co != nil {
			ioutil.WriteFile(currentTokenFile(), []byte(c.ContactToken), 0600)
		}
		ret(co, err)
	},
}

var clientContactAuthCmd = &cobra.Command{
	Use:   "login",
	Short: "Login and keep seesion",
	Run: func(cmd *cobra.Command, args []string) {
		c.ContactEmail = email
		c.ContactPassword = password
		co, err := c.Contact().Authorize()
		if co != nil {
			ioutil.WriteFile(currentTokenFile(), []byte(c.ContactToken), 0600)
		}
		ret(co, err)
	},
}

var clientContactUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update contact",
	Run: func(cmd *cobra.Command, args []string) {
		co, err := c.Contact().Update(name, email, password)
		if co != nil {
			ioutil.WriteFile(currentTokenFile(), []byte(c.ContactToken), 0600)
		}
		ret(co, err)
	},
}

var clientContactGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get contact",
	Run: func(cmd *cobra.Command, args []string) {
		ret(c.Contact().Get(email))
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(clientContactCmd)
	clientCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "contact email")
	clientCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "contact password (min 8 chars)")

	clientContactCmd.AddCommand(clientContactCreateCmd)
	clientContactCreateCmd.Flags().StringVarP(&name, "name", "n", "", "contact name")
	clientContactCmd.AddCommand(clientContactUpdateCmd)
	clientContactUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "contact name")
	clientContactCmd.AddCommand(clientContactGetCmd)
	clientContactCmd.AddCommand(clientContactAuthCmd)
}
