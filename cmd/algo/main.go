package main

import (
	"algo/internal/tui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.New(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
