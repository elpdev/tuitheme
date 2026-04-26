package tuitheme

import "charm.land/lipgloss/v2"

func Phosphor() Theme {
	bright := lipgloss.Color("#9FE8B0")
	muted := lipgloss.Color("#5A8A68")
	subtle := lipgloss.Color("#788E80")
	bg := lipgloss.Color("#111A2C")
	selected := lipgloss.Color("#18233D")
	divider := lipgloss.Color("#2A3752")
	amber := lipgloss.Color("#FFB347")
	cyan := lipgloss.Color("#6FD0E3")

	return Theme{
		Name:       "Phosphor",
		Background: bg,
		Text:       lipgloss.NewStyle().Foreground(bright).Background(bg),
		Muted:      lipgloss.NewStyle().Foreground(muted).Background(bg),
		Title:      lipgloss.NewStyle().Bold(true).Foreground(amber).Background(bg),
		Selected:   lipgloss.NewStyle().Bold(true).Foreground(bright).Background(selected),
		Header:        lipgloss.NewStyle().Foreground(bright).Background(bg).Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(divider).Padding(0, 1),
		HeaderAccent:  lipgloss.NewStyle().Foreground(amber).Background(bg).Bold(true),
		Sidebar:       lipgloss.NewStyle().Background(bg).Border(lipgloss.NormalBorder(), false, true, false, false).BorderForeground(divider).Padding(1, 1),
		Main:          lipgloss.NewStyle().Foreground(bright).Background(bg).Padding(1, 2),
		Footer:        lipgloss.NewStyle().Foreground(subtle).Background(bg).Border(lipgloss.NormalBorder(), true, false, false, false).BorderForeground(divider).Padding(0, 1),
		Modal:         lipgloss.NewStyle().Foreground(bright).Background(bg).Border(lipgloss.RoundedBorder()).BorderForeground(amber).Padding(1, 2),
		PaletteAccent: lipgloss.NewStyle().Foreground(amber).Background(bg).Bold(true),
		Border:        lipgloss.NewStyle().Foreground(divider).Background(bg),
		Info:          lipgloss.NewStyle().Foreground(cyan).Background(bg),
		Warn:          lipgloss.NewStyle().Foreground(amber).Background(bg),
	}
}
