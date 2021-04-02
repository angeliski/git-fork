package git

import (
	"github.com/angeliski/git-fork/executor"
)

type Repository struct {
	path string
	verboseMode bool
}

// TODO is good separate options? Maybe
type PullOptions struct {
	RemoteName string
	Branch string
}

type FetchOptions struct {
	RemoteName string
	Branch string
}

type PushOptions struct {
	RemoteName string
	Branch string
}

type RemoteOptions struct {
	Name string
	Url string
}


func NewRepository(path string, verboseMode bool) (Repository,error) {
	// TODO check if is a git repository
	return Repository{path: path, verboseMode: verboseMode}, nil
}

func (repo Repository) Fetch(options *FetchOptions) error {
	cmdArr := []string{"fetch", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}

func (repo Repository) Pull(options *PullOptions ) error {
	cmdArr := []string{"pull", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}


func (repo Repository) Push(options *PushOptions) error {
	cmdArr := []string{"push", options.RemoteName, options.Branch}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}

func (repo Repository) AddRemote(options *RemoteOptions) error {
	cmdArr := []string{"remote", "add", options.Name, options.Url}

	return executor.RunGitOperation(cmdArr, repo.path, repo.verboseMode)
}