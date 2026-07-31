package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"scalys/internal/theory"
)

func (m model) View() string {
	var b strings.Builder

	title := "Build the scale"
	if m.mode == theory.TriadMode {
		title = "Build the triad"
	}
	b.WriteString(titleStyle.Render("SCALYS — " + title))
	b.WriteString("\n\n")

	label := fmt.Sprintf("%s %s", m.answer[0], m.scaleType)
	if m.mode == theory.TriadMode {
		label += " TRIAD"
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
		b.WriteString(helpStyle.Render("type a-g write note · # sharpen it · enter/→ next · ← back · backspace clear · s submit · m mode · n new · q quit"))
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

	nums := make([]string, len(m.slots))
	for i := range nums {
		nums[i] = lipgloss.NewStyle().Width(8).Align(lipgloss.Center).Render(fmt.Sprint(i + 1))
	}
	numRow := streakStyle.Render(strings.Join(nums, ""))

	return row + "\n" + numRow
}

func (m model) renderPattern() string {
	if m.mode == theory.TriadMode {
		return ""
	}
	syms := m.scaleType.PatternSymbols()[:6]
	return patternStyle.Render("Pattern: " + strings.Join(syms, "  "))
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
