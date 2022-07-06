package cmd

import (
	"video_server/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Long:  "start server",
	Run: func(cmd *cobra.Command, args []string) {
		// 启动项目
		server.Run()
	},
}

func Execute() error {
	return serverCmd.Execute()
}

func init() {
	serverCmd.Flags().BoolP("help", "h", false, "")
}
