package main

import (
	"layzy/src/components"

	"charm.land/bubbles/v2/textinput"
)

type SpaceLst struct {
	cursor int
	list   []Space
}

type Space struct {
	name  string
	tasks []string
}

type model struct {
	space         SpaceLst
	selectedSpace map[int]struct{}
	selectedTask  map[int]struct{}
	inputs        components.Inputs
	isAdding      bool
	isDeleting    bool

	// Meta Details
	width  uint
	height uint
}

func initialModel() model {
	return model{
		spaces: []Space{
			{
				name: "Space 1", tasks: []string{
					"Install Go",
					"Learn Elm Architecture",
				},
			},
			{
				name: "Space 2", tasks: []string{
					"Build Bubble Tea v2 app",
					"Add Lip Gloss styling",
				}},
		},
		selectedSpace: make(map[int]struct{}),
		selectedTask:  make(map[int]struct{}),
		inputs: components.Inputs{
			TextInput: textinput.New(),
		},
	}
}
