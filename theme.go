package tuitheme

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Theme struct {
	Name       string
	Background color.Color

	Text          lipgloss.Style
	Muted         lipgloss.Style
	Title         lipgloss.Style
	Selected      lipgloss.Style
	Header        lipgloss.Style
	HeaderAccent  lipgloss.Style
	Sidebar       lipgloss.Style
	Main          lipgloss.Style
	Footer        lipgloss.Style
	Modal         lipgloss.Style
	PaletteAccent lipgloss.Style
	Border        lipgloss.Style
	Info          lipgloss.Style
	Warn          lipgloss.Style
}

func BuiltIns() []Theme {
	return []Theme{Phosphor()}
}

func ByName(themes []Theme, name string) (Theme, bool) {
	for _, candidate := range themes {
		if candidate.Name == name {
			return candidate, true
		}
	}
	return Theme{}, false
}

func Next(themes []Theme, current string) Theme {
	if len(themes) == 0 {
		return Theme{}
	}
	for i, candidate := range themes {
		if candidate.Name == current {
			return themes[(i+1)%len(themes)]
		}
	}
	return themes[0]
}
