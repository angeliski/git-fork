package git

import "github.com/angeliski/git-fork/executor"

// HasRemote returns if the remoteName is present on the repository
func (repo Repository) HasRemote(remoteName string) bool {

	cmdArr := []string{"remote", "get-url", remoteName}

	err := executor.RunGitOperation(cmdArr, repo.path, false)

	return err == nil
}
