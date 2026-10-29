package main

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = uint(msg.Width)
		m.height = uint(msg.Height)
		return m, nil

	case tea.KeyPressMsg:

		// if m.isAdding {
		// 	switch msg.String() {
		// 	case "enter":
		// 		if m.inputs.TextInput.Value() != "" {
		// 			m.tasks = append(m.tasks, Task{
		// 				title:       m.inputs.TextInput.Value(),
		// 			})
		// 		}
		// 		m.isAdding = false
		// 		m.inputs.TextInput.Blur()
		// 		return m, nil
		// 	case "esc":
		// 		m.isAdding = false
		// 		m.inputs.TextInput.Blur()
		// 		return m, nil
		// 	}

		// 	m.inputs.TextInput, cmd = m.inputs.TextInput.Update(msg)
		// 	return m, cmd
		// }

		// if m.isDeleting {
		// 	switch msg.String() {
		// 	case "enter":
		// 		m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
		// 		m.isDeleting = false

		// 	case "esc":
		// 		m.isDeleting = false
		// 	}

		// 	return m, nil
		// }

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
			// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursors.spaceCursor > 0 {
				m.cursors.spaceCursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursors.spaceCursor < len(m.spaces)-1 {
				m.cursors.spaceCursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selectedSpace[m.cursors.spaceCursor]
			if ok {
				delete(m.selectedSpace, m.cursors.spaceCursor)
			} else {
				m.selectedSpace[m.cursors.spaceCursor] = struct{}{}
			}

		// Add new task
		case "N":
			m.isAdding = true
			m.inputs.TextInput.Focus()
			m.inputs.TextInput.SetValue("")
			return m, textinput.Blink

			// Delete current task
			// case "D":
			// 	if len(m.spaces) > 0 {
			// 		m.isDeleting = true
			// 	}
		}

	}

	if m.isAdding {
		m.inputs.TextInput, cmd = m.inputs.TextInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
