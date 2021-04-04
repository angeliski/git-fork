package executor

import (
	"os"
	"os/exec"
)

type osExecutor struct {
	verboseMode bool
}

// NewOsExecutor returns a exec runs external commands
func NewOsExecutor(verboseMode bool) CommandExecutor {
	return &osExecutor{
		verboseMode: verboseMode,
	}
}

func runLoudly(cmd *exec.Cmd, verboseMode bool) error {
	if verboseMode {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}

	return cmd.Run()
}

// RunGitOperation Run a git command in the path
func (e *osExecutor) RunGitOperation(commands []string, path string) error {
	var args []string
	args = append(args, "git")
	args = append(args, "-C")
	args = append(args, path)
	args = append(args, commands...)

	cmd := exec.Command(args[0], args[1:]...)

	return runLoudly(cmd, e.verboseMode)
}
