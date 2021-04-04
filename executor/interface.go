package executor

// CommandExecutor provides a way to run commands at low level
type CommandExecutor interface {
	RunGitOperation(commands []string, path string) error
}
