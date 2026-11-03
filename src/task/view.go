package task

import (
	"layzy/src/ui"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {

	var taskContent strings.Builder

	taskContent.WriteString(ui.TitleStyle.Render("Tasks"))
	taskContent.WriteString("\n\n")

	if currentTasks := m.Tasks; len(currentTasks) > 0 {
		for i, task := range currentTasks {
			if m.TaskCursor == uint(i) {
				taskContent.WriteString("[>]")
			} else {
				taskContent.WriteString("[ ]")
			}

			taskContent.WriteString(ui.ItemStyle.Render(task.Title))
			taskContent.WriteByte('\n')
		}
	} else {
		taskContent.WriteString("No tasks available.")
	}

	return tea.NewView(taskContent.String())
}
