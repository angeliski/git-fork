package cmd

import (
	"github.com/angeliski/git-fork/app"
	"github.com/angeliski/git-fork/domain/repository"
	"github.com/angeliski/git-fork/executor"
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

	repo, err := git.NewRepositoryService(repositoryPath, executor.NewOsExecutor(verboseMode))
	if err != nil {
		log.Println(err)
		return err
	}

	model := repository.Model{
		RemoteName: "upstream",
		Branch:     "master", //TODO como lidar com master/main?
	}

	service := app.NewBusinessService(repo)

	return service.Sync(model)
}
