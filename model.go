package main

import "charm.land/bubbles/v2/textinput"

type model struct {
	tasks      []Task
	cursor     int
	selected   map[int]struct{}
	textInput  textinput.Model
	isAdding   bool
	isDeleting bool

	// Meta Details
	width  uint
	height uint
}

func initialModel() model {
	newTodo := textinput.New()
	newTodo.Placeholder = "Add new to do"

	return model{
		tasks: []Task{
			{title: "Install Go", description: "Download and install the latest version of Go from golang.org.\nMake sure GOPATH is set."},
			{title: "Learn Elm Architecture", description: "Understand the Model, View, Update cycle.\nThis is crucial for building state machines."},
			{title: "Build Bubble Tea v2 app", description: "Use charm.land/bubbletea/v2 to write the multi-pane TUI."},
			{title: "Add Lip Gloss styling", description: "Use lipgloss.JoinHorizontal to render the layout."},
		},
		selected:  make(map[int]struct{}),
		textInput: newTodo,
	}
}

type Task struct {
	title       string
	description string
}
