package app

import (
	"fmt"
	"layzy/src/ui"
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	// Introduce stack interface for smaller terminal size
	if m.Width < 50 || m.Height < 12 {
		return tea.NewView("Terminal window too small. Please resize.")
	}

	// Background multi pane layout
	availableWidth := m.Width
	spacesPaneWidth := (availableWidth * 3) / 10
	tasksPaneWidth := (availableWidth * 7) / 10
	paneHeight := m.Height

	spacesContent := m.Model.View()
	tasksContent := m.Spaces[m.SpaceCursor].View()

	spacesPaneStyle := ui.PaneStyle(m.Focus == FocusSpaces)
	tasksPaneStyle := ui.PaneStyle(m.Focus == FocusTasks)

	spacesPane := spacesPaneStyle.Width(int(spacesPaneWidth)).Height(int(paneHeight)).Render(spacesContent.Content)
	tasksPane := tasksPaneStyle.Width(int(tasksPaneWidth)).Height(int(paneHeight)).Render(tasksContent.Content)

	finalRender := lipgloss.JoinHorizontal(lipgloss.Top, spacesPane, tasksPane)

	if m.IsAdding {
		log.Printf("Adding New Task")
		m.overlayModal(finalRender)

		return tea.NewView(m.overlayModal(finalRender))
	}

	// if m.IsAdding {
	// 	dimmedBG := ui.DimBackground(backgroundLayout)
	// 	modalHeader := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0055")).Render("Create New Task")
	// 	modalFooter := ui.HelpStyle.Render("enter: save • esc: cancel")

	// 	modalContent := fmt.Sprintf("%s\n\n%s\n\n%s\n\n", modalHeader, m.Inputs.TextInput.View(), modalFooter)
	// 	modalBox := ui.ModalStyle.Render(modalContent)

	// 	modalW := lipgloss.Width(modalBox)
	// 	modalH := lipgloss.Height(modalBox)
	// 	centerX := (int(m.Width) - modalW) / 2
	// 	centerY := (int(m.Height) - modalH) / 2

	// 	finalOverlayView := ui.PlaceOverlay(centerX, centerY, modalBox, dimmedBG)
	// 	return tea.NewView(finalOverlayView)
	// }

	return tea.NewView(finalRender)
}

func (m model) overlayModal(background string) string {
	dimmedBG := ui.DimBackground(background)

	modalHeader := "Create New Task"
	modalFooter := ui.HelpStyle.Render("enter: save • esc: cancel")

	modalContent := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n",
		modalHeader,
		m.Inputs.TextInput.View(),
		modalFooter)
	modalBox := ui.ModalStyle.Render(modalContent)

	centerX := (int(m.Width) - lipgloss.Width(modalBox)) / 2
	centerY := (int(m.Height) - lipgloss.Height(modalBox)) / 2

	return ui.PlaceOverlay(centerX, centerY, modalBox, dimmedBG)
}
