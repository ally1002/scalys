package ui

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"

	"scalys/internal/theory"
)

type gameState int

const (
	statePlaying gameState = iota
	stateSubmitted
)

type model struct {
	root      int
	scaleType theory.ScaleType
	answer    []string
	slots     [7]string
	cursor    int
	state     gameState
	results   [7]bool
	streak    int
}

func New() model {
	m := model{}
	m.newRound()
	return m
}

func (m *model) newRound() {
	prevRoot, prevType := m.root, m.scaleType
	for {
		m.root = rand.Intn(12)
		m.scaleType = theory.ScaleType(rand.Intn(2))
		if m.root != prevRoot || m.scaleType != prevType {
			break
		}
	}
	m.answer = theory.BuildScale(m.root, m.scaleType)
	m.slots = [7]string{}
	m.slots[0] = m.answer[0]
	m.cursor = 1
	m.state = statePlaying
	m.results = [7]bool{}
}

func (m model) filled() bool {
	for _, s := range m.slots {
		if s == "" {
			return false
		}
	}
	return true
}

func (m model) Init() tea.Cmd {
	return nil
}
