package main

import "charm.land/lipgloss/v2"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).PaddingRight(1).
			MarginTop(1).MarginBottom(1).
			Border(lipgloss.RoundedBorder())

	listPanelStyle = lipgloss.NewStyle().
			Width(30).Height(10).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1)

	detailPanelStyle = lipgloss.NewStyle().
				Width(40).Height(10).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#00FFFF")).
				Padding(1)

	modalStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("FF0055")).
			Padding(1)
)
