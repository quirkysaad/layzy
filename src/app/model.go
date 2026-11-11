package app

import (
	"layzy/src/components"
	"layzy/src/space"
	"layzy/src/task"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type FocusState int

const (
	FocusSpaces FocusState = iota
	FocusTasks
)

type model struct {
	space.Model
	Focus      FocusState
	Inputs     components.Inputs
	IsAdding   bool
	IsDeleting bool

	// Meta Details
	Width  uint
	Height uint
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func New() model {
	return initialModel()
}

func initialModel() model {
	return model{
		Spaces: []space.Space{
			{
				Name: "Space 1", Tasks: []task.Task{{
					Title: "Install Go",
					Done:  false,
				}, {
					Title: "Learn Elm Architecture",
					Done:  false,
				}},
			}, {
				Name: "Space 2", Tasks: []task.Task{{Title: "Build Bubble Tea v2 app", Done: false}},
			},
		},
		SpaceCursor: 0,
		Inputs: components.Inputs{
			TextInput: textinput.New(),
		},
		Focus: FocusSpaces,
	}
}
