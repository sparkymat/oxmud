package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func CreateLayout() Layout {
	return Layout{}
}

type Layout struct {
	width  int
	height int
}

func (l Layout) Init() tea.Cmd {
	return nil
}

func (l Layout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedL := l

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		updatedL.width = msg.Width
		updatedL.height = msg.Height
		return updatedL, nil
	}

	return updatedL, nil
}

func (l Layout) View() string {
	return ""
}
