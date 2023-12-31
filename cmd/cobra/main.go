package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/**
 * @title: cobra CLI example
 * @description: a simple example of how to use cobra to create a CLI with subcommands
 * @date: 2023-12-31
 * @usage: go run cmd/cobra/main.go hello alice
 */
func main() {
	Execute()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "an example of how to use cobra to create a CLI with subcommands",
}

var helloCmd = &cobra.Command{
	Use:   "hello [name]",
	Short: "Prints hello to the provided name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s!\n", name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(helloCmd)
}
