package command

import "os/exec"

//go:generate mockgen -destination=mocks/commander_mock.go -package=mocks github.com/omegion/go-command Interface
// Interface is an interface for Command.
type Interface interface {
	Output(string, ...string) ([]byte, error)
}

// Command is custom wrapper for exec.Command.
type Command struct{}

// Output is for executing for exec.Command.
func (c Command) Output(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}
