/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"godemo/internal/goweb/gogin/proxy/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	address string
	port    int
	isDebug bool
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A multi-protocol proxy supporting HTTP and SSH, including interactive SSH sessions.",
	Long: `This is a powerful and flexible proxy server that supports:
  - HTTP Proxy: Forwarding HTTP requests with logging and authentication support.
  - SSH Proxy: Relaying SSH connections securely.
  - Interactive SSH Sessions: Providing real-time SSH session forwarding with enhanced control.
It is designed for developers and sysadmins who need a robust solution for handling HTTP and SSH traffic efficiently.`,

	Run: func(cmd *cobra.Command, args []string) {
		if isDebug {
			gin.SetMode(gin.DebugMode)
		}
		routers.StartProxyService(address, port)
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().IntVarP(&port, "port", "p", 8090, "Port to run the server on.")
	proxyCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "Server address.")
	proxyCmd.Flags().BoolVarP(&isDebug, "debug", "d", false, "Run the server debug mode.")
}
