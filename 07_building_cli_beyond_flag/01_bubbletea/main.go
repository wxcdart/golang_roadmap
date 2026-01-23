package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model int

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m++
		case "down", "j":
			m--
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Counter: %d\n\nControls: up/k increment, down/j decrement, q quit", m)
}

func main() {
	p := tea.NewProgram(model(0))
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
