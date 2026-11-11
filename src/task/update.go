package task

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {
		// The "up" key move the cursor up
		case "up":
			if m.TaskCursor > 0 {
				m.TaskCursor--
			}

		// The "down" key move the cursor down
		case "down":
			if m.TaskCursor < uint(len(m.Tasks)-1) {
				m.TaskCursor++
			}
		}
	}

	return m, nil

}
