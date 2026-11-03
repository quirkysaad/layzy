package ui

import "charm.land/lipgloss/v2"

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).PaddingRight(1).
			MarginTop(1).MarginBottom(1).
			Border(lipgloss.RoundedBorder())

	BasePaneStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#595959")).
			Padding(1)

	FocusedPaneStyle = BasePaneStyle.BorderForeground(lipgloss.Color("#7756e6"))

	ModalStyle = lipgloss.NewStyle().
			Width(80).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF0055")).
			Padding(1, 2)

	SelectedItemStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FFFF")).PaddingLeft(1)
	ItemStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).PaddingLeft(1)
	HelpStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
)

func PaneStyle(focused bool) lipgloss.Style {
	if focused {
		return FocusedPaneStyle
	} else {
		return BasePaneStyle
	}
}
