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

		if m.isAdding {
			switch msg.String() {
			case "enter":
				if m.textInput.Value() != "" {
					m.tasks = append(m.tasks, Task{
						title:       m.textInput.Value(),
						description: "Sample Description",
					})
				}
				m.isAdding = false
				m.textInput.Blur()
				return m, nil
			case "esc":
				m.isAdding = false
				m.isAdding = false
				m.textInput.Blur()
				return m, nil
			}

			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		if m.isDeleting {
			switch msg.String() {
			case "enter":
				m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
				m.isDeleting = false

			case "esc":
				m.isDeleting = false
			}

			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
			// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		// Add new task
		case "N":
			m.isAdding = true
			m.textInput.Focus()
			m.textInput.SetValue("")
			return m, textinput.Blink

		// Delete current task
		case "D":
			if len(m.tasks) > 0 {
				m.isDeleting = true
			}
		}

	}

	if m.isAdding {
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}
