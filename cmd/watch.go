package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fastwand/internal/ui"
	"fastwand/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [directory]",
	Short: "Watch for changes and rebuild CSS",
	Long:  `Start Tailwind in watch mode to automatically rebuild CSS when files change.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}

		// Get absolute path
		absPath, err := filepath.Abs(directory)
		if err != nil {
			fmt.Println(utils.ErrorStyle.Render(err.Error()))
			os.Exit(1)
		}

		// Start spinner
		p := tea.NewProgram(ui.NewSpinner("Starting watch process..."))

		// Start the Python process
		pythonCmd := exec.Command("python", "-c", fmt.Sprintf(`
from fastwand.cli import watch_command
watch_command("%s")
		`, absPath))

		stdout, err := pythonCmd.StdoutPipe()
		if err != nil {
			fmt.Println(utils.ErrorStyle.Render(err.Error()))
			os.Exit(1)
		}

		if err := pythonCmd.Start(); err != nil {
			fmt.Println(utils.ErrorStyle.Render(err.Error()))
			os.Exit(1)
		}

		// Create a scanner to read Python output
		scanner := bufio.NewScanner(stdout)

		// Process status updates from Python
		go func() {
			for scanner.Scan() {
				msg := scanner.Text()
				switch {
				case strings.HasPrefix(msg, "STATUS:"):
					p.Send(strings.TrimPrefix(msg, "STATUS:"))
				case strings.HasPrefix(msg, "ERROR:"):
					p.Send(fmt.Errorf(strings.TrimPrefix(msg, "ERROR:")))
				case strings.HasPrefix(msg, "DONE:"):
					p.Send(true)
				}
			}
		}()

		if _, err := p.Run(); err != nil {
			fmt.Println(utils.ErrorStyle.Render(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
