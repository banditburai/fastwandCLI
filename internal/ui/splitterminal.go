package ui

import (
	"fastwand/internal/process"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WatchModel struct {
	leftView       *TerminalView
	rightView      *TerminalView
	width          int
	height         int
	focused        string // "left" or "right"
	processManager *process.Manager
}

type TerminalView struct {
	viewport viewport.Model
	title    string
	color    string
}

func NewWatchModel() *WatchModel {
	return &WatchModel{
		leftView: &TerminalView{
			viewport: viewport.New(0, 0),
			title:    "Tailwind Watcher",
			color:    "#01FAC6",
		},
		rightView: &TerminalView{
			viewport: viewport.New(0, 0),
			title:    "Python Server",
			color:    "#04B575",
		},
		focused: "left",
	}
}

func (m *WatchModel) Init() tea.Cmd {
	return nil
}

func (m *WatchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Update viewport sizes
		halfWidth := (m.width / 2) - 4
		viewHeight := m.height - 4
		m.leftView.viewport.Width = halfWidth
		m.leftView.viewport.Height = viewHeight
		m.rightView.viewport.Width = halfWidth
		m.rightView.viewport.Height = viewHeight
	case TailwindOutputMsg:
		// Initialize content if empty
		currentContent := m.leftView.viewport.View()
		if currentContent == "" {
			m.leftView.viewport.SetContent(string(msg))
		} else {
			m.leftView.viewport.SetContent(currentContent + "\n" + string(msg))
		}
		// Only scroll if we have content
		if m.leftView.viewport.Height > 0 {
			m.leftView.viewport.GotoBottom()
		}
	case ServerOutputMsg:
		currentContent := m.rightView.viewport.View()
		if currentContent == "" {
			m.rightView.viewport.SetContent(string(msg))
		} else {
			m.rightView.viewport.SetContent(currentContent + "\n" + string(msg))
		}
		if m.rightView.viewport.Height > 0 {
			m.rightView.viewport.GotoBottom()
		}
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	}
	return m, cmd
}

func (m *WatchModel) View() string {
	halfWidth := m.width/2 - 2

	// Create two independent terminal views with their own borders
	leftStyle := lipgloss.NewStyle().
		Width(halfWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(func() string {
			if m.focused == "left" {
				return "#01FAC6"
			}
			return "#666666"
		}())).
		MarginRight(1)

	rightStyle := lipgloss.NewStyle().
		Width(halfWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(func() string {
			if m.focused == "right" {
				return "#04B575"
			}
			return "#666666"
		}()))

	// Create help text
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Align(lipgloss.Center).
		Width(m.width)
	helpText := helpStyle.Render("TAB to switch panes • ↑/↓ to scroll • q to quit")

	// Join horizontally and vertically
	terminals := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftStyle.Render(m.leftView.View(halfWidth-2, m.height-2, m.focused == "left")),
		rightStyle.Render(m.rightView.View(halfWidth-2, m.height-2, m.focused == "right")),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		terminals,
		helpText,
	)
}

type TailwindOutputMsg string
type ServerOutputMsg string

func (m *WatchModel) SetProcessManager(pm *process.Manager) {
	m.processManager = pm
}

func (v *TerminalView) View(width, height int, focused bool) string {
	// Header
	header := lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Bold(true).
		Render(v.title)

	// Set viewport dimensions
	v.viewport.Width = width
	v.viewport.Height = height - 4 // Account for header and padding

	// Let the viewport handle the content rendering
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		v.viewport.View(),
	)
}

func (m *WatchModel) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		if m.processManager != nil {
			// Add a channel to signal cleanup completion
			done := make(chan struct{})
			go func() {
				m.processManager.Cleanup()
				time.Sleep(200 * time.Millisecond)
				close(done)
			}()

			// Wait for cleanup or timeout
			select {
			case <-done:
			case <-time.After(500 * time.Millisecond):
			}
		}
		return m, tea.Quit
	case "tab":
		if m.focused == "left" {
			m.focused = "right"
		} else {
			m.focused = "left"
		}
	case "up", "k":
		if m.focused == "left" {
			m.leftView.viewport.LineUp(1)
		} else {
			m.rightView.viewport.LineUp(1)
		}
	case "down", "j":
		if m.focused == "left" {
			m.leftView.viewport.LineDown(1)
		} else {
			m.rightView.viewport.LineDown(1)
		}
	case "pgup":
		if m.focused == "left" {
			m.leftView.viewport.HalfViewUp()
		} else {
			m.rightView.viewport.HalfViewUp()
		}
	case "pgdown":
		if m.focused == "left" {
			m.leftView.viewport.HalfViewDown()
		} else {
			m.rightView.viewport.HalfViewDown()
		}
	}
	return m, nil
}
