package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"scalys/internal/theory"
)

func (m model) View() string {
	if m.modal {
		return m.renderModal()
	}

	var b strings.Builder

	title := "Build the " + strings.ToLower(m.mode.String())
	b.WriteString(titleStyle.Render("SCALYS — " + title))
	b.WriteString("\n\n")

	label := fmt.Sprintf("%s %s", m.answer[0], m.scaleType)
	if m.mode != theory.ScaleMode {
		label += " " + m.mode.String()
	}
	challenge := challengeStyle.Render("Challenge:  " + label)
	streak := streakStyle.Render(fmt.Sprintf("Streak: %d", m.streak))
	b.WriteString(challenge + "          " + streak)
	b.WriteString("\n\n")

	b.WriteString(m.renderSlots())
	if pattern := m.renderPattern(); pattern != "" {
		b.WriteString("\n" + pattern)
	}
	b.WriteString("\n\n")

	if m.state == statePlaying {
		b.WriteString(helpStyle.Render("type a-g write note · # sharpen it · enter/→ next · ← back · backspace clear · s submit · m mode settings · n new · q quit"))
	} else {
		b.WriteString(m.renderResult())
		b.WriteString("\n\n")
		b.WriteString(helpStyle.Render("enter/n new round · q quit"))
	}
	b.WriteString("\n")

	return lipgloss.NewStyle().Padding(1, 2).Render(b.String())
}

func (m model) renderSlots() string {
	cells := make([]string, len(m.slots))
	for i, note := range m.slots {
		label := note
		if label == "" {
			label = "__"
		}
		style := slotStyle
		switch {
		case m.state == stateSubmitted && m.results[i]:
			style = slotCorrectStyle
		case m.state == stateSubmitted && !m.results[i]:
			style = slotWrongStyle
			label = fmt.Sprintf("%s→%s", note, m.answer[i])
		case i == 0:
			style = slotRootStyle
		case i == m.cursor:
			style = slotFocusedStyle
		}
		cells[i] = style.Render(label)
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, cells...)

	degrees := theory.DegreeNumbers(m.mode, m.scaleType)
	nums := make([]string, len(degrees))
	for i, d := range degrees {
		nums[i] = lipgloss.NewStyle().Width(8).Align(lipgloss.Center).Render(fmt.Sprint(d))
	}
	numRow := streakStyle.Render(strings.Join(nums, ""))

	return row + "\n" + numRow
}

func (m model) renderPattern() string {
	if m.mode != theory.ScaleMode {
		return ""
	}
	syms := m.scaleType.PatternSymbols()[:6]
	return patternStyle.Render("Pattern: " + strings.Join(syms, "  "))
}

func (m model) renderModal() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("SCALYS — Mode Settings"))
	b.WriteString("\n\n")

	rows := []struct {
		mode  theory.PracticeMode
		key   string
		label string
	}{
		{theory.ScaleMode, "S", "Scale"},
		{theory.TriadMode, "T", "Triad"},
		{theory.PentatonicMode, "P", "Pentatonic"},
	}
	for _, r := range rows {
		line := fmt.Sprintf("[%s] %s", r.key, r.label)
		if r.mode == m.mode {
			line += " ✓"
		}
		style := helpStyle
		if r.mode == m.modalRow {
			style = challengeStyle
			line = "› " + line
		} else {
			line = "  " + line
		}
		b.WriteString(style.Render(line))
		b.WriteString("\n")
	}
	b.WriteString("\n")

	qs := make([]string, 3)
	for i, q := range []quality{qualityRandom, qualityMajor, qualityMinor} {
		label := q.String()
		if q == m.quality {
			label = challengeStyle.Render(label)
		} else {
			label = helpStyle.Render(label)
		}
		qs[i] = label
	}
	b.WriteString(strings.Join(qs, helpStyle.Render(" | ")))
	b.WriteString("\n\n")

	b.WriteString(helpStyle.Render("s/t/p select mode · ↑/↓ highlight · enter apply · ←/→ change quality · esc close"))
	b.WriteString("\n")

	return lipgloss.NewStyle().Padding(1, 2).Render(b.String())
}

func (m model) renderResult() string {
	correct := 0
	for _, r := range m.results {
		if r {
			correct++
		}
	}
	if correct == len(m.results) {
		return correctMsgStyle.Render("✓ Correct! Well done.")
	}
	return wrongMsgStyle.Render(fmt.Sprintf("✗ %d of %d right — wrong slots show yours→expected.", correct-1, len(m.results)-1))
}
