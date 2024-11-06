package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"fastwand/internal/ui"
	"fastwand/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the FastHTML app with minified CSS",
	Run: func(cmd *cobra.Command, args []string) {
		// Create and start spinner program
		p := tea.NewProgram(ui.NewSpinner("Starting FastHTML app..."))

		// Create error channel to track overall success
		errChan := make(chan error, 1)

		// Set up signal handling
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		// Run tasks in goroutine
		go func() {
			// Get Tailwind path
			p.Send("Locating Tailwind...")
			tailwindPath, err := utils.GetTailwindPath(".")
			if err != nil {
				errChan <- fmt.Errorf("tailwind binary not found\ncheck directory and run 'fastwand init' first")
				p.Send(fmt.Errorf("tailwind binary not found\ncheck directory and run 'fastwand init' first"))
				return
			}

			// Delete existing output.css if it exists
			outputPath := filepath.Join(".", "static", "css", "output.css")
			if err := os.Remove(outputPath); err != nil && !os.IsNotExist(err) {
				errChan <- err
				p.Send(fmt.Errorf("failed to remove old CSS: %w", err))
				return
			}

			// Build CSS
			p.Send("Building CSS...")
			tailwindCmd := exec.Command(tailwindPath,
				"-i", "static/css/input.css",
				"-o", "static/css/output.css",
				"--minify")
			if err := tailwindCmd.Run(); err != nil {
				errChan <- err
				p.Send(fmt.Errorf("failed to build CSS: %w", err))
				return
			}

			close(errChan)
			p.Send(true)
		}()

		// Run the spinner program
		if _, err := p.Run(); err != nil {
			fmt.Println(utils.ErrorStyle.Render("Error: " + err.Error()))
			return
		}

		// Check if any errors occurred during setup
		select {
		case err := <-errChan:
			if err != nil {
				return
			}
		default:
		}

		// Start server
		serverCmd := exec.Command("python", "main.py")
		serverCmd.Stdout = os.Stdout
		serverCmd.Stderr = os.Stderr
		serverCmd.Stdin = os.Stdin

		// Handle graceful shutdown
		go func() {
			<-sigChan
			fmt.Print("\n")
			if serverCmd.Process != nil {
				// Send interrupt signal instead of kill
				serverCmd.Process.Signal(os.Interrupt)

				// Wait for server to shutdown gracefully
				done := make(chan struct{})
				go func() {
					serverCmd.Wait()
					close(done)
				}()

				// Wait for graceful shutdown or timeout
				select {
				case <-done:
					fmt.Print("\n")
				case <-time.After(2 * time.Second):
					// Force kill if taking too long
					serverCmd.Process.Kill()
				}
			}
			time.Sleep(200 * time.Millisecond)
			os.Exit(0)
		}()

		// Start server and only print link if successful
		if err := serverCmd.Start(); err != nil {
			fmt.Println(utils.ErrorStyle.Render(fmt.Sprintf("server error: %v", err)))
			return
		}

		// Wait a short moment to check if server starts successfully
		time.Sleep(500 * time.Millisecond)

		// Only show link if we got here with no errors and server is running
		if serverCmd.Process != nil && serverCmd.ProcessState == nil {
			fmt.Println(utils.InfoStyle.Render("http://localhost:5001"))
		}

		if err := serverCmd.Wait(); err != nil {
			if _, ok := err.(*exec.ExitError); !ok {
				fmt.Println(utils.ErrorStyle.Render(fmt.Sprintf("server error: %v", err)))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
