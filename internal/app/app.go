package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const title = `
   ___       __  __ _   _ ____  
  / _ \__  _|  \/  | | | |  _ \ 
 | | | \ \/ / |\/| | | | | | | |
 | |_| |>  <| |  | | |_| | |_| |
  \___//_/\_\_|  |_|\___/|____/ 
                                
`

type App struct {
	Width  int
	Height int
}

func Create() App {
	return App{}
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		updatedA := a
		updatedA.Width = msg.Width
		updatedA.Height = msg.Height
		return updatedA, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	}

	return a, nil
}

func (a App) View() string {
	style := lipgloss.NewStyle().
		Width(40).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#99ccaa")).
		BorderBackground(lipgloss.Color("#222222")).
		Foreground(lipgloss.Color("#99ccaa")).
		Background(lipgloss.Color("#222222")).
		PaddingBottom(1).
		PaddingTop(1).
		PaddingLeft(4).
		PaddingRight(4)

	return lipgloss.Place(
		a.Width,
		a.Height,
		lipgloss.Center,
		lipgloss.Center,
		style.Render(title),
		lipgloss.WithWhitespaceBackground(lipgloss.Color("#222222")),
	)
}
