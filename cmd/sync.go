package cmd

import (
	"errors"
	"fmt"
	"github.com/angeliski/git-fork/git"
	"github.com/spf13/cobra"
	"log"
)

var sync = &cobra.Command{
	Use:   "sync",
	Short: "used to run a fetch/pull operation in your repository",
	RunE:  syncRun,
}

func init() {
	// a branch que eu quero fazer o sync
	// o nome do remote que eu quero usar
	// desabilitar o push padrão

	getRootCmd().AddCommand(sync)
}

func syncRun(cmd *cobra.Command, args []string) error {
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

	if !repo.HasRemote("upstream") {
		return errors.New("remote upstream not configured. Use git-fork fork` to configure")
	}

	fmt.Println("Fetching")
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "upstream",
		Branch:     "master", //TODO como lidar com master/main?
	})

	if err != nil {
		return err
	}

	fmt.Println("Pulling")
	err = repo.Pull(&git.PullOptions{
		RemoteName: "upstream",
		Branch:     "master", //TODO como lidar com master/main?
	})

	if err != nil {
		return err
	}

	fmt.Println("Pushing")
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Branch:     "master", //TODO como lidar com master/main?
	})

	return err
}
