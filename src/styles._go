package main

import "charm.land/lipgloss/v2"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).PaddingRight(1).
			MarginTop(1).MarginBottom(1).
			Border(lipgloss.RoundedBorder())

	basePaneStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#595959")).
			Padding(1)

	focusedPaneStyle = basePaneStyle.BorderForeground(lipgloss.Color("#7756e6"))

	modalStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
			Width(80).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("FF0055")).
			Padding(1, 2)

	selectedItemStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FFFF")).PaddingLeft(1)
	itemStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).PaddingLeft(1)
	helpStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).MarginTop(1)
)
