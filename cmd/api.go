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
	"github.com/mirrorhub-io/platform/controllers"
	"github.com/mirrorhub-io/platform/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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

func init() {
	RootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&addr, "addr", "l", "127.0.0.1", "Bind addr.")
	apiCmd.Flags().StringVarP(&port, "port", "p", "", "Bind port.")
	RootCmd.AddCommand(gatewayCmd)
	gatewayCmd.Flags().StringVarP(&addr, "addr", "l", "127.0.0.1", "Bind addr.")
	gatewayCmd.Flags().StringVarP(&port, "port", "p", "", "Bind port.")
	gatewayCmd.Flags().StringVarP(&api, "api", "a", "127.0.0.1:9000", "Grpc server.")
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
