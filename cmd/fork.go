package cmd

import (
	"github.com/angeliski/git-fork/app"
	"github.com/angeliski/git-fork/domain/repository"
	"github.com/angeliski/git-fork/executor"
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
	getRootCmd().AddCommand(fork)
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

	repo, err := git.NewRepositoryService(repositoryPath, executor.NewOsExecutor(verboseMode))
	if err != nil {
		log.Println(err)
		return err
	}

	service := app.NewBusinessService(repo)

	return service.Fork(
		repository.Model{
			RemoteName: "upstream",
			URLRemote:  args[0],
		})
}
