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
	"io/ioutil"
	"os/user"
	"strconv"

	"github.com/mirrorhub-io/platform/client"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var clientMirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Mirrorhub mirror utils",
}

var clientContactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Mirrorhub contacts utils",
}

var clientServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Mirrorhub service utils",
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

var clientMirrorListCmd = &cobra.Command{
	Use:   "list",
	Short: "List mirrors",
	Run: func(cmd *cobra.Command, args []string) {
		ret(c.Mirror().List())
	},
}

var mirrorConnectTo string

var clientMirrorConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect mirror and mirroring endpoint together",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			exit("Expected MirrorID")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			exit(err.Error())
		}
		endpoint_id, err := strconv.Atoi(mirrorConnectTo)
		if err != nil {
			exit(err.Error())
		}
		ret(c.Mirror().Connect(int32(i), int32(endpoint_id)))
	},
}

var clientMirrorGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Find mirror by id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			exit("Expected MirrorID")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			exit(err.Error())
		}
		ret(c.Mirror().FindById(int32(i)))
	},
}

var (
	m_bandwidth  string
	m_domain     string
	m_ipv4       string
	m_ipv6       string
	m_name       string
	m_storage    string
	m_traffic    string
	m_service_id string
)

var clientMirrorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update mirror by id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			exit("Expected MirrorID")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			exit(err.Error())
		}
		ret(c.Mirror().UpdateById(int32(i), mirrorFromFlags()))
	},
}

var clientMirrorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Register new mirror",
	Run: func(cmd *cobra.Command, args []string) {
		ret(c.Mirror().Create(mirrorFromFlags()))
	},
}

func mirrorFromFlags() *pb.Mirror {
	m := &pb.Mirror{}
	if m_domain != "" {
		m.Domain = m_domain
	}
	if m_ipv4 != "" {
		m.Ipv4 = m_ipv4
	}
	if m_ipv6 != "" {
		m.Ipv6 = m_ipv6
	}
	if m_name != "" {
		m.Name = m_name
	}
	if m_storage != "" {
		i, _ := strconv.Atoi(m_storage)
		m.Storage = int64(i)
	}
	if m_traffic != "" {
		i, _ := strconv.Atoi(m_traffic)
		m.Traffic = int64(i)
	}
	if m_service_id != "" {
		i, _ := strconv.Atoi(m_service_id)
		m.Service = &pb.Service{
			Id: int32(i),
		}
	}
	if m_bandwidth != "" {
		i, _ := strconv.Atoi(m_bandwidth)
		m.Bandwidth = int64(i)
	}
	return m
}

var clientServiceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List services",
	Run: func(cmd *cobra.Command, args []string) {
		ret(c.Service().List())
	},
}

var (
	s_name    string
	s_storage string
)

func serviceFromFlags(files []string) *pb.Service {
	s := &pb.Service{}
	if s_name != "" {
		s.Name = s_name
	}
	if s_storage != "" {
		i, _ := strconv.Atoi(s_storage)
		s.Storage = int64(i)
	}
	s.Files = files
	return s
}

var clientServiceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create service",
	Run: func(cmd *cobra.Command, args []string) {
		ret(c.Service().Create(serviceFromFlags(args)))
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(clientContactCmd)
	clientCmd.AddCommand(clientMirrorCmd)
	clientCmd.AddCommand(clientServiceCmd)
	clientCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "contact email")
	clientCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "contact password (min 8 chars)")

	clientContactCmd.AddCommand(clientContactCreateCmd)
	clientContactCreateCmd.Flags().StringVarP(&name, "name", "n", "", "contact name")
	clientContactCmd.AddCommand(clientContactUpdateCmd)
	clientContactUpdateCmd.Flags().StringVarP(&name, "name", "n", "", "contact name")
	clientContactCmd.AddCommand(clientContactGetCmd)
	clientContactCmd.AddCommand(clientContactAuthCmd)

	clientMirrorCmd.AddCommand(clientMirrorListCmd)
	clientMirrorCmd.AddCommand(clientMirrorGetCmd)
	clientMirrorCmd.AddCommand(clientMirrorConnectCmd)
	clientMirrorConnectCmd.Flags().StringVarP(&mirrorConnectTo, "endpoint-id", "e", "", "The service you want to connect to")

	clientMirrorCmd.AddCommand(clientMirrorUpdateCmd)
	clientMirrorUpdateCmd.Flags().StringVarP(&m_ipv4, "ipv4", "4", "", "IPv4 address")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_ipv6, "ipv6", "6", "", "IPv6 address")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_name, "name", "n", "", "Mirror's display name")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_domain, "domain", "d", "", "Mirror's domain name")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_storage, "storage", "s", "", "Storage usage limit")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_bandwidth, "bandwidth", "b", "", "Bandwith limit")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_traffic, "traffic", "t", "", "Monthly traffic limit")
	clientMirrorUpdateCmd.Flags().StringVarP(&m_service_id, "service-id", "", "", "Service id")

	clientMirrorCmd.AddCommand(clientMirrorCreateCmd)
	clientMirrorCreateCmd.Flags().StringVarP(&m_ipv4, "ipv4", "4", "", "IPv4 address")
	clientMirrorCreateCmd.Flags().StringVarP(&m_ipv6, "ipv6", "6", "", "IPv6 address")
	clientMirrorCreateCmd.Flags().StringVarP(&m_name, "name", "n", "", "Mirror's display name")
	clientMirrorCreateCmd.Flags().StringVarP(&m_domain, "domain", "d", "", "Mirror's domain name")
	clientMirrorCreateCmd.Flags().StringVarP(&m_storage, "storage", "s", "", "Storage usage limit")
	clientMirrorCreateCmd.Flags().StringVarP(&m_bandwidth, "bandwidth", "b", "", "Bandwith limit")
	clientMirrorCreateCmd.Flags().StringVarP(&m_traffic, "traffic", "t", "", "Monthly traffic limit")
	clientMirrorCreateCmd.Flags().StringVarP(&m_service_id, "service-id", "", "", "Service id")

	clientServiceCmd.AddCommand(clientServiceListCmd)
	clientServiceCmd.AddCommand(clientServiceCreateCmd)
	clientServiceCreateCmd.Flags().StringVarP(&s_name, "name", "n", "", "Service's display name")
	clientServiceCreateCmd.Flags().StringVarP(&s_storage, "storage", "s", "", "Storage requirement")
}
