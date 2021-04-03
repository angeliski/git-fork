package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// GetPath returns the path to run the commands
func GetPath(cmd *cobra.Command) (string, error) {
	flag := cmd.Flag("path")
	if flag.Value.String() != "" {
		return flag.Value.String(), nil
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return path, nil
}

// IsVerboseMode returns if verbose mode is enabled
func IsVerboseMode(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("verbose")
}
