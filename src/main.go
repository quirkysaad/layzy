package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
