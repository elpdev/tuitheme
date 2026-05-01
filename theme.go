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
	Accent        lipgloss.Style
	Title         lipgloss.Style
	Selected      lipgloss.Style
	Disabled      lipgloss.Style
	Header        lipgloss.Style
	HeaderAccent  lipgloss.Style
	Sidebar       lipgloss.Style
	Main          lipgloss.Style
	Footer        lipgloss.Style
	Modal         lipgloss.Style
	PaletteAccent lipgloss.Style
	Border        lipgloss.Style
	Focus         lipgloss.Style
	Info          lipgloss.Style
	Success       lipgloss.Style
	Warn          lipgloss.Style
	Error         lipgloss.Style
}

func BuiltIns() []Theme {
	return []Theme{
		Phosphor(),
		Dracula(),
		TokyoNight(),
		TokyoNightStorm(),
		CatppuccinMocha(),
		CatppuccinMacchiato(),
		Nord(),
		GruvboxDark(),
		GruvboxLight(),
		Kanagawa(),
		RosePine(),
		RosePineMoon(),
		SolarizedDark(),
		SolarizedLight(),
		OneDark(),
		EverforestDark(),
		Monokai(),
	}
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
