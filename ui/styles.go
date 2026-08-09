package ui

import "github.com/charmbracelet/lipgloss"

var (
	Border = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#040037")).
		Bold(true)

	Fih = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#2b0000"))

	Bubble = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#2c3133"))

	Sea = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0059fd"))

	Music = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#990052")).
		Italic(true)

	Text = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#694a00")).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#B78CFF")).
		Padding(0, 2)
)

