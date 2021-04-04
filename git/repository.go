package git

import (
	"github.com/angeliski/git-fork/domain/repository"
	"github.com/angeliski/git-fork/executor"
)

// Repository is a representation from git operations
type Repository struct {
	path     string
	executor executor.CommandExecutor
}

// NewRepositoryService return the struct to run git operations
func NewRepositoryService(path string, executor executor.CommandExecutor) (repository.Service, error) {
	// TODO check if is a git repository
	return Repository{
		path:     path,
		executor: executor,
	}, nil
}

// Fetch the repository based in options
func (r Repository) Fetch(repository repository.Model) error {
	cmdArr := []string{"fetch", repository.RemoteName, repository.Branch}

	return r.executor.RunGitOperation(cmdArr, r.path)
}

// Pull the repository based in options
func (r Repository) Pull(repository repository.Model) error {
	cmdArr := []string{"pull", repository.RemoteName, repository.Branch}

	return r.executor.RunGitOperation(cmdArr, r.path)
}

// Push the repository based in options
func (r Repository) Push(repository repository.Model) error {
	cmdArr := []string{"push", repository.RemoteName, repository.Branch}

	return r.executor.RunGitOperation(cmdArr, r.path)
}

// AddRemote to the repository based in options
func (r Repository) AddRemote(repository repository.Model) error {
	cmdArr := []string{"remote", "add", repository.RemoteName, repository.URLRemote}

	return r.executor.RunGitOperation(cmdArr, r.path)
}

// HasRemote returns if the remoteName is present on the repository
func (r Repository) HasRemote(remoteName string) bool {

	cmdArr := []string{"remote", "get-url", remoteName}

	err := r.executor.RunGitOperation(cmdArr, r.path)

	return err == nil
}
