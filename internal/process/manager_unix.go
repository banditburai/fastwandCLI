//go:build !windows

package process

import (
	"os/exec"

	"golang.org/x/sys/unix"
)

func (pm *Manager) SetupCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &unix.SysProcAttr{Setpgid: true}
}

func (pm *Manager) KillProcess(cmd *exec.Cmd) error {
	if cmd != nil && cmd.Process != nil {
		if pgid, err := unix.Getpgid(cmd.Process.Pid); err == nil {
			return unix.Kill(-pgid, unix.SIGKILL)
		}
		return cmd.Process.Kill()
	}
	return nil
}
