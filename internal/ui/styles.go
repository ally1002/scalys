package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("212"))

	challengeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("220"))

	streakStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	slotStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1).
			Width(6).
			Align(lipgloss.Center).
			BorderForeground(lipgloss.Color("240"))

	slotFocusedStyle = slotStyle.
				BorderForeground(lipgloss.Color("212")).
				Foreground(lipgloss.Color("212")).
				Bold(true)

	slotRootStyle = slotStyle.
			BorderForeground(lipgloss.Color("220")).
			Foreground(lipgloss.Color("220"))

	slotCorrectStyle = slotStyle.
				BorderForeground(lipgloss.Color("42")).
				Foreground(lipgloss.Color("42")).
				Bold(true)

	slotWrongStyle = slotStyle.
			BorderForeground(lipgloss.Color("196")).
			Foreground(lipgloss.Color("196")).
			Bold(true)

	patternStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	correctMsgStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42"))

	wrongMsgStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
