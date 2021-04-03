package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "git-fork",
	Short: "git-fork helps you to maintain you repository updated",
}

var repositoryPath string
var verbose bool

func init() {
	//cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&repositoryPath, "path", "", "Target directory to run git operations")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Run operations with verbose mode")
}

// getRootCmd returns the rootCmd
func getRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
