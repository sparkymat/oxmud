package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
}

func CreateApp() App {
	return App{}
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	}

	return a, nil
}

func (a App) View() string {
	return "this is a thing"
}

func main() {
	p := tea.NewProgram(CreateApp())

	if _, err := p.Run(); err != nil {
		fmt.Printf("failed with err: %s", err.Error())
		os.Exit(1)
	}
}
