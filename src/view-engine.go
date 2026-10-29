package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {

	// Introduce stack interface for smaller terminal size
	if m.width < 50 || m.height < 12 {
		return tea.NewView("Terminal window too small. Please resize.")
	}

	// Background multi pane layout
	availableWidth := m.width - 4
	spacesPaneWidth := (availableWidth * 3) / 10
	tasksPaneWidth := (availableWidth * 7) / 10
	paneHeight := m.height - 4

	// Spaces pane
	var spacesContent strings.Builder

	spacesContent.WriteString(titleStyle.Render("Spaces"))
	spacesContent.WriteString("\n\n")

	for i, space := range m.spaces {
		row := fmt.Sprintf("%s", space.name)

		if m.cursors.spaceCursor == i {
			spacesContent.WriteByte('>')
		} else {
			spacesContent.WriteByte(' ')
		}
		spacesContent.WriteString(itemStyle.Render(row))
		spacesContent.WriteByte('\n')

	}

	spacesPane := spacesPaneStyle.Width(int(spacesPaneWidth)).Height(int(paneHeight)).Render(spacesContent.String())

	// Tasks pane
	var tasksContent string
	// if len(m.spaces) > 0 {
	// 	activeTask := m.spaces[m.cursor]
	// 	tasksContent = titleStyle.Background(lipgloss.Color("#00FFFF")).Foreground(lipgloss.Color("#000000")).Render("Tasks") + "\n\n"
	// 	tasksContent += lipgloss.NewStyle().Bold(true).Render(activeTask.title) + "\n\n"
	// 	tasksContent += lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")).Render(activeTask.description)
	// } else {
	// 	tasksContent = "No tasks available."
	// }

	tasksContent = "Yet to be developed"
	detailPane := detailPaneStyle.Width(int(tasksPaneWidth)).Height(int(paneHeight)).Render(tasksContent)

	// Combine into final background layout
	backgroundLayout := lipgloss.JoinHorizontal(lipgloss.Top, spacesPane, detailPane)
	backgroundLayout += "\n" + helpStyle.Render("space: toggle • n: new task • q: quit • ↑/↓: navigate")

	if m.isAdding {
		dimmedBG := dimBackground(backgroundLayout)
		modalHeader := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0055")).Render("Create New Task")
		modalFooter := helpStyle.Render("enter: save • esc: cancel")

		modalContent := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", modalHeader, m.inputs.TextInput.View(), modalFooter)
		modalBox := modalStyle.Render(modalContent)

		modalW := lipgloss.Width(modalBox)
		modalH := lipgloss.Height(modalBox)
		centerX := (int(m.width) - modalW) / 2
		centerY := (int(m.height) - modalH) / 2

		finalOverlayView := placeOverlay(centerX, centerY, modalBox, dimmedBG)
		return tea.NewView(finalOverlayView)
	}

	// if m.isDeleting {
	// 	s = fmt.Sprintf("Are you sure you want to Delete : \n\n%s\n\n", m.tasks[m.cursor])
	// 	return tea.NewView(s)
	// }

	// 	for i, task := range m.tasks {
	// 		cursor := " "

	// 		if m.cursor == i {
	// 			cursor = ">"
	// 		}

	// 		checked := " "
	// 		if _, ok := m.selected[i]; ok {
	// 			checked = "✔"
	// 		}

	// 		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.title)
	// 	}

	// s += "\nThanks\n"

	return tea.NewView(backgroundLayout)
}
