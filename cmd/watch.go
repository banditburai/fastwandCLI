package cmd

import (
	"fastwand/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [directory]",
	Short: "Start Tailwind watch mode for development",
	Long: `Start Tailwind CSS in watch mode for development.
This will watch your input CSS file and rebuild on changes.
Run 'python main.py' in a separate terminal to start the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory := "."
		if len(args) > 0 {
			directory = args[0]
		}

		// Check if tailwindcss exists
		tailwindPath := filepath.Join(directory, "tailwindcss")
		if runtime.GOOS == "windows" {
			tailwindPath += ".exe"
		}
		if _, err := os.Stat(tailwindPath); os.IsNotExist(err) {
			utils.ErrorStyle.Render("Tailwind executable not found. Did you run 'fastwand init' first?")
			return
		}

		utils.InfoStyle.Render("Starting Tailwind watch mode...")
		utils.InfoStyle.Render("\nNOTE: Run 'python main.py' in a separate terminal")

		// Start watch process
		watchCmd := exec.Command(tailwindPath,
			"-i", "assets/input.css",
			"-o", "assets/output.css",
			"--watch")
		watchCmd.Dir = directory
		watchCmd.Stdout = os.Stdout
		watchCmd.Stderr = os.Stderr

		if err := watchCmd.Run(); err != nil {
			utils.ErrorStyle.Render("Error running Tailwind: " + err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
