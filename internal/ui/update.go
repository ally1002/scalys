package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "n":
		m.newRound()
		return m, nil
	case "m":
		m.toggleMode()
		return m, nil
	}

	if m.state == stateSubmitted {
		if keyMsg.String() == "enter" {
			m.newRound()
		}
		return m, nil
	}

	switch keyMsg.String() {
	case "left":
		if m.cursor > 1 {
			m.cursor--
		}
	case "right", "enter":
		if m.cursor < len(m.slots)-1 {
			m.cursor++
		}
	case "a", "b", "c", "d", "e", "f", "g":
		m.slots[m.cursor] = strings.ToUpper(keyMsg.String())
	case "#":
		m.toggleSharp()
	case "backspace":
		if m.slots[m.cursor] != "" {
			m.slots[m.cursor] = ""
		} else if m.cursor > 1 {
			m.cursor--
			m.slots[m.cursor] = ""
		}
	case "s":
		if m.filled() {
			m.submit()
		}
	}
	return m, nil
}

// toggleSharp sharpens the focused note and advances to the next slot;
// on an already-sharp note it removes the sharp and stays.
// E and B have no sharp in the sharps-only note set, so they are left as-is.
func (m *model) toggleSharp() {
	note := m.slots[m.cursor]
	if note == "" {
		return
	}
	switch {
	case strings.HasSuffix(note, "#"):
		m.slots[m.cursor] = strings.TrimSuffix(note, "#")
	case note != "E" && note != "B":
		m.slots[m.cursor] = note + "#"
		if m.cursor < len(m.slots)-1 {
			m.cursor++
		}
	}
}

func (m *model) submit() {
	m.state = stateSubmitted
	allCorrect := true
	for i := range m.slots {
		m.results[i] = m.slots[i] == m.answer[i]
		if !m.results[i] {
			allCorrect = false
		}
	}
	if allCorrect {
		m.streak++
	} else {
		m.streak = 0
	}
}
