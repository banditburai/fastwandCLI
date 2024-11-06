//go:build windows

package process

import (
	"os/exec"
)

func (pm *Manager) SetupCmd(cmd *exec.Cmd) {
	// Windows doesn't need special process group handling
}

func (pm *Manager) KillProcess(cmd *exec.Cmd) error {
	if cmd != nil && cmd.Process != nil {
		return cmd.Process.Kill()
	}
	return nil
}
