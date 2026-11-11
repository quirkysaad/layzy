package space

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
			if m.SpaceCursor > 0 {
				m.SpaceCursor--
			}

		// The "down" key move the cursor down
		case "down":
			if m.SpaceCursor < uint(len(m.Spaces)-1) {
				m.SpaceCursor++
			}
		}
	}
	return m, nil

}
