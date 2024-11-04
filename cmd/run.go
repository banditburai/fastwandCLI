package cmd

import (
	"fastwand/internal/utils"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [directory]",
	Short: "Build CSS and run the server",
	Long: `Build minified CSS with Tailwind and start the Python server.
This command is meant for production use.`,
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

		// Build CSS
		utils.InfoStyle.Render("Building CSS...")
		buildCmd := exec.Command(tailwindPath,
			"-i", "assets/input.css",
			"-o", "assets/output.css",
			"--minify")
		buildCmd.Dir = directory
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			utils.ErrorStyle.Render("Error building CSS: " + err.Error())
			return
		}

		utils.SuccessStyle.Render("CSS built successfully!")

		// Start Python server
		utils.InfoStyle.Render("Starting server...")
		serverCmd := exec.Command("python", "main.py")
		serverCmd.Dir = directory
		serverCmd.Stdout = os.Stdout
		serverCmd.Stderr = os.Stderr

		// Set up signal handling
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		if err := serverCmd.Start(); err != nil {
			utils.ErrorStyle.Render("Error starting server: " + err.Error())
			return
		}

		// Wait for either server completion or interrupt
		go func() {
			<-sigChan
			utils.InfoStyle.Render("\nShutting down server...")
			serverCmd.Process.Signal(os.Interrupt)
		}()

		if err := serverCmd.Wait(); err != nil {
			if err.Error() != "signal: interrupt" {
				utils.ErrorStyle.Render("Error running server: " + err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
