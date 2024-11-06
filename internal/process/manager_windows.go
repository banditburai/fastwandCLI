//go:build windows

package process

import (
	"os/exec"
	"syscall"
)

func (pm *Manager) SetupCmd(cmd *exec.Cmd) {
	// Windows doesn't need special process group handling
}

func (pm *Manager) GetProcessGroupID(cmd *exec.Cmd) int {
	// Windows doesn't support process groups in the same way
	return 0
}

func (pm *Manager) TerminateProcessGroup(pgid int) error {
	if cmd := pm.PythonCmd; cmd != nil && cmd.Process != nil {
		cmd.Process.Signal(syscall.SIGTERM)
	}
	if cmd := pm.TailwindCmd; cmd != nil && cmd.Process != nil {
		cmd.Process.Signal(syscall.SIGTERM)
	}
	return nil
}

func (pm *Manager) KillProcessGroup(pgid int) error {
	if cmd := pm.PythonCmd; cmd != nil && cmd.Process != nil {
		cmd.Process.Kill()
	}
	if cmd := pm.TailwindCmd; cmd != nil && cmd.Process != nil {
		cmd.Process.Kill()
	}
	return nil
}
