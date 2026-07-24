package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("SCALYS — Build the scale"))
	b.WriteString("\n\n")

	challenge := challengeStyle.Render(fmt.Sprintf("Challenge:  %s %s", m.answer[0], m.scaleType))
	streak := streakStyle.Render(fmt.Sprintf("Streak: %d", m.streak))
	b.WriteString(challenge + "          " + streak)
	b.WriteString("\n\n")

	b.WriteString(m.renderSlots())
	b.WriteString("\n")
	b.WriteString(m.renderPattern())
	b.WriteString("\n\n")

	if m.state == statePlaying {
		b.WriteString(helpStyle.Render("type a-g write note · # sharpen it · enter/→ next · ← back · backspace clear · s submit · n new · q quit"))
	} else {
		b.WriteString(m.renderResult())
		b.WriteString("\n\n")
		b.WriteString(helpStyle.Render("enter/n new round · q quit"))
	}
	b.WriteString("\n")

	return lipgloss.NewStyle().Padding(1, 2).Render(b.String())
}

func (m model) renderSlots() string {
	cells := make([]string, 7)
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

	nums := make([]string, 7)
	for i := range nums {
		nums[i] = lipgloss.NewStyle().Width(8).Align(lipgloss.Center).Render(fmt.Sprint(i + 1))
	}
	numRow := streakStyle.Render(strings.Join(nums, ""))

	return row + "\n" + numRow
}

func (m model) renderPattern() string {
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
	if correct == 7 {
		return correctMsgStyle.Render("✓ Correct! Well done.")
	}
	return wrongMsgStyle.Render(fmt.Sprintf("✗ %d of 6 right — wrong slots show yours→expected.", correct-1))
}
