package space

import (
	"layzy/src/task"

	tea "charm.land/bubbletea/v2"
)

type Space struct {
	Name string
	task.Model
}

type Model struct {
	Spaces      []Space
	SpaceCursor uint
}

func (m Model) Init() tea.Cmd {
	return nil
}
