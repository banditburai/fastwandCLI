package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"syscall"

	"fastwand/internal/process"
	"fastwand/internal/ui"
	"fastwand/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func runTailwind(directory, tailwindPath string, outputChan chan<- ui.TailwindOutputMsg, done <-chan struct{}, pm *process.Manager) {
	outputChan <- ui.TailwindOutputMsg("Starting Tailwind watcher...")

	// Delete existing output.css if it exists
	outputPath := filepath.Join(directory, "static", "css", "output.css")
	if err := os.Remove(outputPath); err != nil && !os.IsNotExist(err) {
		outputChan <- ui.TailwindOutputMsg(fmt.Sprintf("Failed to remove old CSS: %v", err))
		return
	}

	// Start the watch process
	watchCmd := exec.Command(tailwindPath,
		"-i", "static/css/input.css",
		"-o", "static/css/output.css",
		"--watch")
	watchCmd.Dir = directory

	// Create a pipe for stdin
	stdin, err := watchCmd.StdinPipe()
	if err != nil {
		outputChan <- ui.TailwindOutputMsg(fmt.Sprintf("Failed to create stdin pipe: %v", err))
		return
	}

	// Connect the command's stdout and stderr directly
	watchCmd.Stdout = writeWrapper{outputChan, directory}
	watchCmd.Stderr = writeWrapper{outputChan, directory}

	// Start the command
	pm.SetupCmd(watchCmd)

	if err := watchCmd.Start(); err != nil {
		outputChan <- ui.TailwindOutputMsg(fmt.Sprintf("Watch process error: %v", err))
		return
	}

	// Keep stdin open and write periodically to keep watch alive
	go func() {
		defer stdin.Close()
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				stdin.Write([]byte{'\n'})
			}
		}
	}()

	// Wait for process to finish
	if err := watchCmd.Wait(); err != nil {
		outputChan <- ui.TailwindOutputMsg(fmt.Sprintf("Watch process error: %v", err))
	}
}

// Helper type to wrap our channel as an io.Writer
type writeWrapper struct {
	ch        chan<- ui.TailwindOutputMsg
	directory string
}

func (w writeWrapper) Write(p []byte) (n int, err error) {
	lines := strings.Split(strings.TrimSpace(string(p)), "\n")
	for _, line := range lines {
		if line != "" {
			w.ch <- ui.TailwindOutputMsg(line)
		}
	}
	return len(p), nil
}

func runServer(directory string, outputChan chan<- ui.ServerOutputMsg, done <-chan struct{}, pm *process.Manager) {
	cmd := exec.Command("python", "-c", `
import uvicorn
import main
uvicorn.run(
	"main:app",
	host="0.0.0.0",
	port=5001,
	reload=True,
	reload_includes=["*.py", "static/css/output.css"]
)`)
	cmd.Dir = directory

	pm.SetupCmd(cmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		outputChan <- ui.ServerOutputMsg(fmt.Sprintf("Error creating pipe: %v", err))
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		outputChan <- ui.ServerOutputMsg(fmt.Sprintf("Error creating pipe: %v", err))
		return
	}

	if err := cmd.Start(); err != nil {
		outputChan <- ui.ServerOutputMsg(fmt.Sprintf("Failed to start server: %v", err))
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			select {
			case <-done:
				return
			case outputChan <- ui.ServerOutputMsg(scanner.Text()):
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			select {
			case <-done:
				return
			case outputChan <- ui.ServerOutputMsg(scanner.Text()):
			}
		}
	}()

	<-done
	cmd.Process.Kill()
	cmd.Wait()
}

var watchCmd = &cobra.Command{
	Use:   "watch [directory]",
	Short: "Start development server with live reload",
	Run: func(cmd *cobra.Command, args []string) {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}

		tailwindPath, err := utils.GetTailwindPath(directory)
		if err != nil {
			fmt.Println(utils.ErrorStyle.Render("tailwind binary not found\ncheck directory and run 'fastwand init' first"))
			return
		}

		tailwindChan := make(chan ui.TailwindOutputMsg)
		serverChan := make(chan ui.ServerOutputMsg)

		pm := process.NewManager()

		model := ui.NewWatchModel()
		model.SetProcessManager(pm)

		p := tea.NewProgram(
			model,
			tea.WithAltScreen(),
		)

		go runTailwind(directory, tailwindPath, tailwindChan, pm.Done, pm)
		go runServer(directory, serverChan, pm.Done, pm)

		go func() {
			for {
				select {
				case msg := <-tailwindChan:
					p.Send(msg)
				case msg := <-serverChan:
					p.Send(msg)
				case <-pm.Done:
					return
				}
			}
		}()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-sigChan
			pm.Cleanup()
			p.Quit()
		}()

		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
