package app

import (
	"layzy/src/space"
	"layzy/src/task"

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

			switch msg.String() {
			case "enter":
				if m.Inputs.TextInput.Value() != "" {
					currentSpace := &m.Spaces[m.SpaceCursor]
					currentSpace.Tasks = append(currentSpace.Tasks, task.Task{Title: m.Inputs.TextInput.Value(), Done: false})
				}
				m.IsAdding = false
				m.Inputs.TextInput.Blur()
				return m, nil
			case "esc":
				m.IsAdding = false
				m.Inputs.TextInput.Blur()
				return m, nil
			}

			m.Inputs.TextInput, cmd = m.Inputs.TextInput.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		// Select Spaces pane
		case "1", "left":
			if m.Focus == FocusTasks {
				m.Focus = FocusSpaces
			}
		// Select Tasks pane
		case "2", "right":
			if m.Focus == FocusSpaces {
				m.Focus = FocusTasks
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

		// Forward remaining keys to the focused pane
		if m.Focus == FocusSpaces {
			updated, cmd := m.Model.Update(msg)
			m.Model = updated.(space.Model)
			return m, cmd
		}

		if m.Focus == FocusTasks {
			currentSpace := &m.Model.Spaces[m.SpaceCursor]
			updated, cmd := currentSpace.Model.Update(msg)
			currentSpace.Model = updated.(task.Model)
			return m, cmd
		}
	}
	return m, nil
}
