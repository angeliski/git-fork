package git

import "github.com/angeliski/git-fork/executor"

func (repo Repository) HasRemote(remoteName string) bool {

	cmdArr := []string{"remote", "get-url", remoteName}

	err := executor.RunGitOperation(cmdArr, repo.path, false)

	return err == nil
}
