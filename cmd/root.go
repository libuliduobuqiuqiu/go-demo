/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   string
	buildTime string
)

var rootCmd = &cobra.Command{
	Use: "godemo",
	Short: `A Go project for learning Go language basics, 
	popular libraries, algorithm implementations, real-world use cases, and related technical documents.`,
	Long: `go-demo is a Go project designed for learning and practicing Go programming. 
	It covers fundamental syntax, the usage of popular libraries, algorithm implementations, real-world business scenarios, and relevant technical documentation. 
	The project serves as a hands-on learning resource for Go developers to explore various aspects of the language and its ecosystem.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of the godemo project.",
	Long:  "Prints the version of this godemo project, including build information such as commit hash and build date if available",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nBuild Time: %s\n", version, buildTime)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
