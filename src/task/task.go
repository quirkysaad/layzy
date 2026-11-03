package task

import (
	tea "charm.land/bubbletea/v2"
)

type Task struct {
	Title string
	Done  bool
}

type Model struct {
	Tasks      []Task
	TaskCursor uint
}

func (m Model) Init() tea.Cmd {
	return nil
}
