package app

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = uint(msg.Width)
		m.Height = uint(msg.Height)
		return m, nil

	case tea.KeyPressMsg:
		// Handle key presses within modal
		if m.IsAdding {
			m.Inputs.TextInput, cmd = m.Inputs.TextInput.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		// Select Spaces pane
		case "1", "left":
			if m.Focus == FocusSpaces {
				m.Focus = FocusTasks
			}
		// Select Tasks pane
		case "2", "right":
			if m.Focus == FocusTasks {
				m.Focus = FocusSpaces
			}
		// Quit Layzy
		case "ctrl+c", "q":
			return m, tea.Quit

		// Add new task
		case "N":
			m.IsAdding = true
			m.Inputs.TextInput.Focus()
			m.Inputs.TextInput.SetValue("")
			return m, textinput.Blink
		}

	}
	return m, nil
}
