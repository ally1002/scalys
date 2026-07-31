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

// quality constrains which scale type a practice mode deals: always major,
// always minor, or a random pick each round.
type quality int

const (
	qualityMajor quality = iota
	qualityMinor
	qualityRandom
)

func (q quality) String() string {
	switch q {
	case qualityMajor:
		return "Major"
	case qualityMinor:
		return "Minor"
	default:
		return "Random"
	}
}

func (q quality) next() quality { return (q + 1) % 3 }
func (q quality) prev() quality { return (q + 2) % 3 }

type model struct {
	root      int
	scaleType theory.ScaleType
	mode      theory.PracticeMode
	quality   quality
	modal     bool
	modalRow  theory.PracticeMode
	answer    []string
	slots     []string
	cursor    int
	state     gameState
	results   []bool
	streak    int
}

func New() model {
	m := model{quality: qualityRandom}
	m.newRound()
	return m
}

func (m *model) newRound() {
	prevRoot, prevType := m.root, m.scaleType
	for {
		m.root = rand.Intn(12)
		switch m.quality {
		case qualityMajor:
			m.scaleType = theory.Major
		case qualityMinor:
			m.scaleType = theory.Minor
		default:
			m.scaleType = theory.ScaleType(rand.Intn(2))
		}
		if m.root != prevRoot || m.scaleType != prevType {
			break
		}
	}
	m.answer = theory.Build(m.root, m.scaleType, m.mode)
	m.slots = make([]string, len(m.answer))
	m.slots[0] = m.answer[0]
	m.cursor = 1
	m.state = statePlaying
	m.results = make([]bool, len(m.answer))
}

// openModal opens the mode-settings modal, starting the highlight on the
// current mode.
func (m *model) openModal() {
	m.modal = true
	m.modalRow = m.mode
}

// selectMode switches practice mode, closes the modal, and deals a new
// round. Used by the s/t/p shortcuts, which pick and enter a mode in one
// step.
func (m *model) selectMode(mode theory.PracticeMode) {
	m.mode = mode
	m.modalRow = mode
	m.modal = false
	m.newRound()
}

// moveModalRow shifts which mode is highlighted in the modal without
// applying it — used by up/down, which just browse.
func (m *model) moveModalRow(delta int) {
	if delta < 0 {
		m.modalRow = (m.modalRow + 2) % 3
	} else {
		m.modalRow = (m.modalRow + 1) % 3
	}
}

// applyModalRow switches practice mode to the highlighted row, closes the
// modal, and deals a new round — used by enter to confirm an up/down pick.
func (m *model) applyModalRow() {
	m.selectMode(m.modalRow)
}

// cycleQuality changes the scale quality for future rounds. It does not
// close the modal or deal a new round — only s/t/p (selectMode) does that.
func (m *model) cycleQuality(delta int) {
	if delta < 0 {
		m.quality = m.quality.prev()
	} else {
		m.quality = m.quality.next()
	}
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
