package process

import (
	"os/exec"
	"syscall"
	"time"
)

type Manager struct {
	PythonCmd   *exec.Cmd
	TailwindCmd *exec.Cmd
	Done        chan struct{}
	Port        int
}

func NewManager() *Manager {
	return &Manager{
		Done: make(chan struct{}),
		Port: 5001, // Default port
	}
}

func (pm *Manager) Cleanup() {
	close(pm.Done)

	// Helper function to gracefully terminate a process
	cleanup := func(cmd *exec.Cmd) {
		if cmd != nil && cmd.Process != nil {
			// Get process group ID before sending signals
			pgid, err := syscall.Getpgid(cmd.Process.Pid)
			if err == nil {
				// Send SIGTERM to process group
				syscall.Kill(-pgid, syscall.SIGTERM)
			}

			// Give it a moment to cleanup
			done := make(chan error)
			go func() {
				done <- cmd.Wait()
			}()

			// Wait for process to exit or timeout
			select {
			case <-time.After(1000 * time.Millisecond):
				// Force kill the process group if still running
				if pgid != 0 {
					syscall.Kill(-pgid, syscall.SIGKILL)
				}
				cmd.Process.Kill()
				// Wait for the process to actually terminate
				cmd.Wait()
			case <-done:
				// Process exited cleanly
			}
		}
	}

	// Clean up both processes with longer timeout
	cleanup(pm.PythonCmd)
	cleanup(pm.TailwindCmd)

	// Additional sleep to ensure sockets are released
	time.Sleep(100 * time.Millisecond)
}

func (pm *Manager) SetPythonCmd(cmd *exec.Cmd) {
	pm.PythonCmd = cmd
}

func (pm *Manager) SetTailwindCmd(cmd *exec.Cmd) {
	pm.TailwindCmd = cmd
}
