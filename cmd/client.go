/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"godemo/internal/goweb/gogin/proxy/client"

	"github.com/spf13/cobra"
)

var wsServer string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Launch internal service test clients.",
	Long: `
	This command is designed to start client modules, such as HTTP、SSH and WebSocket clients,
	which are primarily used for testing and debugging internal service code.
	It allows to verify API connectivity, inspect response, and access service stability,
	making it ideal for development and testing environment.
	`,
}

var webSocketCmd = &cobra.Command{
	Use:   "ws",
	Short: "Start a websocket client.",
	Long: `
	Established a websocket connection to test real-time communication with a server
	`,
	Run: func(cmd *cobra.Command, args []string) {
		client.CommitDeviceTerminalReq(wsServer)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(webSocketCmd)
	webSocketCmd.Flags().StringVarP(&wsServer, "address", "a", "127.0.0.1:8090", "Websocket service address.")
}
