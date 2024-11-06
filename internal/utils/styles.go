package utils

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Title style for headers
	TitleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#01FAC6")).
			Foreground(lipgloss.Color("#030303")).
			Bold(true).
			Padding(0, 1).
			Align(lipgloss.Center)

	// Success messages
	SuccessStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#04B575")).
			Foreground(lipgloss.Color("#04B575")).
			Margin(1).
			Padding(1)

	// Info messages and prompts
	InfoStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#01FAC6")).
			Margin(1).
			Padding(1, 2)

	// Error messages
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF616E")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF616E")).
			Margin(1).
			Padding(0, 1).
			Width(50)

	// Selected item style
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#01FAC6")).
			Bold(true)

	// Description style
	DescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#40BDA3")).
			Align(lipgloss.Center)

	// Help text style
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))

	// URL hint style
	URLHintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#01FAC6")).
			Padding(0, 1).
			Width(50).
			Align(lipgloss.Center)
)
