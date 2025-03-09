/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "godemo",
	Short: `A Go project for learning Go language basics, 
	popular libraries, algorithm implementations, real-world use cases, and related technical documents.`,
	Long: `go-demo is a Go project designed for learning and practicing Go programming. 
	It covers fundamental syntax, the usage of popular libraries, algorithm implementations, real-world business scenarios, and relevant technical documentation. 
	The project serves as a hands-on learning resource for Go developers to explore various aspects of the language and its ecosystem.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
