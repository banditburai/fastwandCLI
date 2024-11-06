//go:build !windows

package process

import (
	"os/exec"

	"golang.org/x/sys/unix"
)

func (pm *Manager) SetupCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &unix.SysProcAttr{Setpgid: true}
}

func (pm *Manager) GetProcessGroupID(cmd *exec.Cmd) int {
	if cmd != nil && cmd.Process != nil {
		if pgid, err := unix.Getpgid(cmd.Process.Pid); err == nil {
			return pgid
		}
	}
	return 0
}

func (pm *Manager) TerminateProcessGroup(pgid int) error {
	if pgid != 0 {
		return unix.Kill(-pgid, unix.SIGTERM)
	}
	return nil
}

func (pm *Manager) KillProcessGroup(pgid int) error {
	if pgid != 0 {
		return unix.Kill(-pgid, unix.SIGKILL)
	}
	return nil
}
