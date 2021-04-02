package git

import (
	"github.com/angeliski/git-fork/executor"
)

// Repository is a representation from git operations
type Repository struct {
	path        string
	verboseMode bool
}

// PullOptions to run the command
type PullOptions struct {
	RemoteName string
	Branch     string
}

// FetchOptions to run the command
type FetchOptions struct {
	RemoteName string
	Branch     string
}

// PushOptions to run the command
type PushOptions struct {
	RemoteName string
	Branch     string
}

// RemoteOptions to run the command
type RemoteOptions struct {
	Name string
	URL  string
}

// NewRepository return the struct to run git operations
func NewRepository(path string, verboseMode bool) (Repository, error) {
	// TODO check if is a git repository
	return Repository{path: path, verboseMode: verboseMode}, nil
}

// Fetch the repository based in options
func (repo Repository) Fetch(options *FetchOptions) error {
	cmdArr := []string{"fetch", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}

// Pull the repository based in options
func (repo Repository) Pull(options *PullOptions) error {
	cmdArr := []string{"pull", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}

// Push the repository based in options
func (repo Repository) Push(options *PushOptions) error {
	cmdArr := []string{"push", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}

// AddRemote to the repository based in options
func (repo Repository) AddRemote(options *RemoteOptions) error {
	cmdArr := []string{"remote", "add", options.Name, options.URL}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}
