package space

import (
	"layzy/src/ui"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// Spaces pane
	var spaceContent strings.Builder

	spaceContent.WriteString(ui.TitleStyle.Render("Spaces"))
	spaceContent.WriteString("\n\n")

	for i, space := range m.Spaces {
		if m.SpaceCursor == uint(i) {
			spaceContent.WriteByte('>')
		} else {
			spaceContent.WriteByte(' ')
		}
		spaceContent.WriteString(ui.ItemStyle.Render(space.Name))
		spaceContent.WriteByte('\n')
	}
	return tea.NewView(spaceContent.String())
}
