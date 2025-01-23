package main

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea/v2"
)

type model struct {
	quitting   bool
	suspending bool
}

func (m model) Init() (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.ResumeMsg:
		m.suspending = false
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c":
			m.quitting = true
			return m, tea.Interrupt
		case "ctrl+z":
			m.suspending = true
			return m, tea.Suspend
		}
	}
	return m, nil
}

func (m model) View() fmt.Stringer {
	if m.suspending || m.quitting {
		return tea.NewFrame("")
	}

	return tea.NewFrame("\nPress ctrl-z to suspend, ctrl+c to interrupt, q, or esc to exit\n")
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("Error running program:", err)
		if errors.Is(err, tea.ErrInterrupted) {
			os.Exit(130)
		}
		os.Exit(1)
	}
}
