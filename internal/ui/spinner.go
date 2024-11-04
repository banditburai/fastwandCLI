package ui

import (
	"fastwand/internal/utils"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type SpinnerModel struct {
	spinner spinner.Model
	message string
	err     error
	done    bool
	success bool
}

func NewSpinner(message string) *SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = utils.SelectedStyle
	return &SpinnerModel{
		spinner: s,
		message: message,
	}
}

func (m *SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, tea.Quit
	case bool:
		m.done = true
		m.success = msg
		return m, tea.Quit
	case string:
		m.message = msg
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *SpinnerModel) View() string {
	if m.err != nil {
		return utils.ErrorStyle.Render(m.err.Error())
	}
	if m.done {
		if m.success {
			return utils.SuccessStyle.Render("🪄 Operation completed successfully!")
		}
		return utils.ErrorStyle.Render("😢 Operation failed!")
	}
	return fmt.Sprintf("%s %s", m.spinner.View(), m.message)
}

func (m *SpinnerModel) SetMessage(message string) {
	m.message = message
}

func (m *SpinnerModel) SetError(err error) {
	m.err = err
}

func (m *SpinnerModel) SetDone(success bool) {
	m.done = true
	m.success = success
}
