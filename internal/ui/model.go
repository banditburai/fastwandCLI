package ui

import (
	"fastwand/internal/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	choices  []string
	cursor   int
	selected string
	width    int
	height   int
}

func InitialModel() Model {
	return Model{
		choices: []string{
			"🌼 Tailwind CSS with DaisyUI",
			"🔷 Vanilla Tailwind CSS",
		},
		width:  0,
		height: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#01FAC6")).
		Padding(1, 2).
		Width(m.width - 4).
		Align(lipgloss.Center)

	content := utils.TitleStyle.Render("Welcome to FastWand🪄") + "\n\n"
	content += utils.DescStyle.Render("FastHTML + Tailwind made easy!") + "\n\n"
	content += "Choose your UI framework:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = utils.SelectedStyle.Render(">")
		}

		choiceStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
		if m.cursor == i {
			choiceStyle = utils.SelectedStyle
		}

		content += fmt.Sprintf("%s %s\n", cursor, choiceStyle.Render(choice))
	}

	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	content += "\n" + helpStyle.Render("(↑/↓ to move, enter to select, q to quit)")

	return boxStyle.Render(content)
}

func (m Model) Selected() string {
	return m.selected
}

var _ tea.Model = (*Model)(nil)
