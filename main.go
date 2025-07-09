package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sparkymat/oxmud/internal/app"
)

func main() {
	p := tea.NewProgram(app.Create(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("failed with err: %s", err.Error())
		os.Exit(1)
	}
}
