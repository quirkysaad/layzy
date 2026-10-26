package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	// s := titleStyle.Render("Organize your day Layzy Way") + "\n"
	if m.width == 0 {
		return tea.NewView("Initializing...")
	}

	// Background multi pane layout
	availableWidth := m.width - 4
	listWidth := (availableWidth * 4) / 10
	detailWidth := (availableWidth * 6) / 10
	paneHeight := m.height - 4

	// List pane
	listContent := titleStyle.Render("Tasks") + "\n\n"
	for i, task := range m.tasks {
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		row := fmt.Sprintf("[%s] %s", checked, task.title)

		if m.cursor == i {
			listContent += ">" + selectedItemStyle.Render(row) + "\n"
		} else {
			listContent += " " + itemStyle.Render(row) + "\n"
		}
	}
	listPane := listPaneStyle.Width(int(listWidth)).Height(int(paneHeight)).Render(listContent)
	// Details pane
	var detailContent string
	if len(m.tasks) > 0 {
		activeTask := m.tasks[m.cursor]
		detailContent = titleStyle.Background(lipgloss.Color("#00FFFF")).Foreground(lipgloss.Color("#000000")).Render("Details") + "\n\n"
		detailContent += lipgloss.NewStyle().Bold(true).Render(activeTask.title) + "\n\n"
		detailContent += lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")).Render(activeTask.description)
	} else {
		detailContent = "No tasks available."
	}
	detailPane := detailPaneStyle.Width(int(detailWidth)).Height(int(paneHeight)).Render(detailContent)

	// Combine into final background layout
	backgroundLayout := lipgloss.JoinHorizontal(lipgloss.Top, listPane, detailPane)
	backgroundLayout += "\n" + helpStyle.Render("space: toggle • n: new task • q: quit • ↑/↓: navigate")

	if m.isAdding {
		modalContent := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0055")).Render("Add New Task") + "\n\n"
		modalContent = m.textInput.View() + "\n\n"
		modalContent += helpStyle.Render("enter: save • esc: cancel")

		modalBox := modalStyle.Background(lipgloss.Color("#111111")).Render(modalContent)

		overlayedView := lipgloss.Place(int(m.width), int(m.height), lipgloss.Center, lipgloss.Center, modalBox, lipgloss.WithWhitespaceChars(backgroundLayout))
		return tea.NewView(overlayedView)
	}

	// if m.isDeleting {
	// 	s = fmt.Sprintf("Are you sure you want to Delete : \n\n%s\n\n", m.tasks[m.cursor])
	// 	return tea.NewView(s)
	// }

	// 	for i, task := range m.tasks {
	// 		cursor := " "
	//
	// 		if m.cursor == i {
	// 			cursor = ">"
	// 		}
	//
	// 		checked := " "
	// 		if _, ok := m.selected[i]; ok {
	// 			checked = "✔"
	// 		}
	//
	// 		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.title)
	// 	}

	// s += "\nThanks\n"

	return tea.NewView(backgroundLayout)
}
