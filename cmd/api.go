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
	"io/ioutil"
	"os"
	"strings"

	"github.com/mirrorhub-io/platform/balancer"
	"github.com/mirrorhub-io/platform/client"
	"github.com/mirrorhub-io/platform/controllers"
	"github.com/mirrorhub-io/platform/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addr string
	port string
	api  string
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start mirrorhub api",
	Run: func(cmd *cobra.Command, args []string) {
		parseApiCommands(viper.GetString("API.base"))
		controllers.StartApi(addr, port)
		defer models.Connection().Close()
	},
}

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start mirrorhub rest-gateway",
	Run: func(cmd *cobra.Command, args []string) {
		parseApiCommands(viper.GetString("API.gateway"))
		if api == "" {
			api = viper.GetString("API.base")
			fmt.Println("Fallback to configured grpc server.")
			if api == "" {
				fmt.Println("Missing grpc server configuration")
				os.Exit(1)
			}
		}
		controllers.StartGateway(addr, port, api)
		defer models.Connection().Close()
	},
}

var (
	lb_mon string
)

var lbCmd = &cobra.Command{
	Use:   "balancer",
	Short: "Start loadbalancer and monitor",
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
		c.ContactEmail = email
		c.ContactPassword = password
	},
	Run: func(cmd *cobra.Command, args []string) {
		if lb_mon == "true" {
			m := balancer.NewMonitor(c)
			m.Preload()
		}
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&addr, "addr", "l", "127.0.0.1", "Bind addr.")
	apiCmd.Flags().StringVarP(&port, "port", "p", "", "Bind port.")

	RootCmd.AddCommand(gatewayCmd)
	gatewayCmd.Flags().StringVarP(&addr, "addr", "l", "127.0.0.1", "Bind addr.")
	gatewayCmd.Flags().StringVarP(&port, "port", "p", "", "Bind port.")
	gatewayCmd.Flags().StringVarP(&api, "api", "a", "127.0.0.1:9000", "Grpc server.")

	RootCmd.AddCommand(lbCmd)
	lbCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "contact email")
	lbCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "contact password (min 8 chars)")
	lbCmd.Flags().StringVarP(&lb_mon, "monitor", "m", "true", "Enable / disable service monitor")
}

func parseApiCommands(config string) {
	a := strings.Split(config, ":")
	if len(a) > 1 {
		if addr == "" {
			addr = a[0]
		}
		if port == "" {
			port = a[1]
		}
	} else {
		fmt.Println("Ignoring config (<addr>:<port> required)")
	}
}
