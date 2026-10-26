package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	s := titleStyle.Render("Organize your day Layzy Way") + "\n"

	if m.isAdding {
		modalContent := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("FF0055")).Render("Add New Task") + "\n\n"
		modalContent = m.textInput.View() + "\n\n"

		renderedModal := modalStyle.Render(modalContent)
		centeredModal := lipgloss.Place(int(m.width), int(m.height), lipgloss.Center, lipgloss.Center, renderedModal)
		return tea.NewView(centeredModal)
	}

	if m.isDeleting {
		s = fmt.Sprintf("Are you sure you want to Delete : \n\n%s\n\n", m.tasks[m.cursor])
		return tea.NewView(s)
	}

	for i, task := range m.tasks {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "✔"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.title)
	}

	s += "\nThanks\n"

	return tea.NewView(s)
}
