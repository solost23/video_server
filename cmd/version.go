package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of server",
	Long:  "version of server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("server version: %v \n", version)
	},
}

func init() {
	serverCmd.AddCommand(versionCmd)
}
