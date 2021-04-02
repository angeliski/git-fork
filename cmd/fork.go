package cmd

import (
	"github.com/angeliski/git-fork/git"
	"github.com/spf13/cobra"
	"log"
)

var fork = &cobra.Command{
	Use:   "fork",
	Short: "used to fork a repository",
	RunE:  forkRun,
}

func init() {
	GetRootCmd().AddCommand(fork)
}

func forkRun(cmd *cobra.Command, args []string) error {
	repositoryPath, err := GetPath(cmd)
	if err != nil {
		return err
	}

	verboseMode, err := IsVerboseMode(cmd)

	if err != nil {
		return err
	}


	repo, err := git.NewRepository(repositoryPath, verboseMode)
	if err != nil {
		log.Println(err)
		return err
	}


	return repo.AddRemote(&git.RemoteOptions{
		Name: "upstream",
		Url:  args[0],
	})
}