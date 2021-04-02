package executor

import (
	"os"
	"os/exec"
)

func runLoudly(cmd *exec.Cmd, verboseMode bool) error {
	if verboseMode {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
	}

	return cmd.Run()
}


func RunGitOperation(commands []string, path string, verboseMode bool) error {
	var args []string
	args = append(args, "git")
	args = append(args, "-C")
	args = append(args,path)
	args = append(args, commands...)

	cmd := exec.Command(args[0],args[1:]...)

	return runLoudly(cmd, verboseMode)
}


