package cmd

import (
	"fastwand/internal/download"
	"fastwand/internal/templates"
	"fastwand/internal/ui"
	"fastwand/internal/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize a new FastHTML + Tailwind project",
	Long: `Initialize a new project with FastHTML and Tailwind CSS.
If no directory is specified, the current directory will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}

		// Create and run our Bubble Tea program
		p := tea.NewProgram(
			ui.InitialModel(),
			tea.WithAltScreen(),       // use alternate screen buffer
			tea.WithMouseCellMotion(), // enable mouse support
		)
		model, err := p.Run()
		if err != nil {
			utils.ErrorStyle.Render("Error running program: " + err.Error())
			return
		}

		// Type assert our model to access its fields
		m, ok := model.(ui.Model)
		if !ok {
			utils.ErrorStyle.Render("Could not get program state")
			return
		}

		// Use the selected framework to initialize the project
		if m.Selected() != "" {
			initProject(directory, m.Selected())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initProject(directory string, selectedFramework string) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		fmt.Println(utils.ErrorStyle.Render("Failed to create directory: " + err.Error()))
		return
	}

	useDaisy := selectedFramework == "🌼 Tailwind CSS with DaisyUI"

	// Create and start spinner program
	p := tea.NewProgram(ui.NewSpinner("Setting up your FastHTML project..."))

	// Run installation tasks in goroutine
	go func() {
		// Download Tailwind
		p.Send("Downloading Tailwind...")
		tailwindPath, err := download.DownloadTailwind(directory, useDaisy)
		if err != nil {
			p.Send(err)
			return
		}

		// Create project files
		p.Send("Creating project files...")
		if err := templates.CreateProjectFiles(directory, useDaisy); err != nil {
			p.Send(err)
			return
		}

		// Initialize Tailwind
		p.Send("Initializing Tailwind...")
		// Use absolute path and ensure file exists
		absPath, err := filepath.Abs(tailwindPath)
		if err != nil {
			p.Send(fmt.Errorf("failed to get absolute path: %w", err))
			return
		}

		// Check if file exists and is executable
		if _, err := os.Stat(absPath); err != nil {
			p.Send(fmt.Errorf("tailwind binary not found: %w", err))
			return
		}

		initCmd := exec.Command(absPath, "init")
		initCmd.Dir = directory
		if err := initCmd.Run(); err != nil {
			p.Send(fmt.Errorf("failed to initialize tailwind: %w", err))
			return
		}

		// Signal completion
		p.Send(true)
	}()

	// Run the spinner program
	if _, err := p.Run(); err != nil {
		fmt.Println(utils.ErrorStyle.Render("Error: " + err.Error()))
		return
	}

	// Show next steps
	fmt.Println(utils.InfoStyle.Render("\nNext steps:" +
		"\n1. cd " + directory +
		"\n2. fastwand watch - Watch for CSS changes (needs two terminals to run app)" +
		"\n3. fastwand run   - Minify CSS and serve app"))
}
