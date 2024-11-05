package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fastwand/internal/utils"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [directory]",
	Short: "Start development server with live reload",
	Long:  `Start Tailwind in watch mode and run the Python development server.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			switch {
			case strings.HasPrefix(msg, "STATUS:"):
				fmt.Println(utils.InfoStyle.Render(strings.TrimPrefix(msg, "STATUS:")))
			case strings.HasPrefix(msg, "ERROR:"):
				fmt.Println(utils.ErrorStyle.Render(strings.TrimPrefix(msg, "ERROR:")))
			case strings.HasPrefix(msg, "DONE:"):
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
