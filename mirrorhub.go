package main

import (
	"fmt"
	"github.com/mirrorhub-io/platform/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
