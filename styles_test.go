package tuitheme

import (
	"fmt"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestBuiltIns(t *testing.T) {
	themes := BuiltIns()
	if len(themes) == 0 {
		t.Fatal("expected built-in themes")
	}

	seen := make(map[string]bool, len(themes))
	for _, theme := range themes {
		if theme.Name == "" {
			t.Fatal("expected built-in theme to define a name")
		}
		if theme.Background == nil {
			t.Fatalf("expected %s to define a background", theme.Name)
		}
		if seen[theme.Name] {
			t.Fatalf("duplicate built-in theme name %q", theme.Name)
		}
		seen[theme.Name] = true
	}

	for _, name := range []string{"Phosphor", "Dracula", "Tokyo Night", "Catppuccin Mocha", "Nord", "Gruvbox Dark"} {
		if !seen[name] {
			t.Fatalf("expected built-ins to include %s", name)
		}
	}
}

func TestByName(t *testing.T) {
	theme, ok := ByName(BuiltIns(), "Phosphor")
	if !ok {
		t.Fatal("expected to find Phosphor")
	}
	if theme.Name != "Phosphor" {
		t.Fatalf("expected Phosphor, got %q", theme.Name)
	}

	if _, ok := ByName(BuiltIns(), "Missing"); ok {
		t.Fatal("expected missing theme not to be found")
	}

	theme, ok = ByName(BuiltIns(), "Dracula")
	if !ok {
		t.Fatal("expected to find Dracula")
	}
	if theme.Name != "Dracula" {
		t.Fatalf("expected Dracula, got %q", theme.Name)
	}
}

func TestNext(t *testing.T) {
	themes := []Theme{{Name: "One"}, {Name: "Two"}}

	if next := Next(themes, "One"); next.Name != "Two" {
		t.Fatalf("expected Two after One, got %q", next.Name)
	}
	if next := Next(themes, "Two"); next.Name != "One" {
		t.Fatalf("expected One after Two, got %q", next.Name)
	}
	if next := Next(themes, "Missing"); next.Name != "One" {
		t.Fatalf("expected fallback One, got %q", next.Name)
	}
	if next := Next(nil, "Missing"); next.Name != "" {
		t.Fatalf("expected zero theme, got %q", next.Name)
	}
}

func TestPhosphorDefinesBackground(t *testing.T) {
	if theme := Phosphor(); theme.Background == nil {
		t.Fatal("expected Phosphor to define a background")
	}
}

func TestPhosphorKeepsChromeAccent(t *testing.T) {
	theme := Phosphor()
	want := lipgloss.Color("#FFB347")

	if got := theme.Title.GetForeground(); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("expected Phosphor Title foreground %v, got %v", want, got)
	}
	if got := theme.HeaderAccent.GetForeground(); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("expected Phosphor HeaderAccent foreground %v, got %v", want, got)
	}
}

func TestThemeBrandAccents(t *testing.T) {
	cases := map[string]string{
		"Phosphor":             "#FFB347",
		"Dracula":              "#BD93F9",
		"Tokyo Night":          "#7AA2F7",
		"Tokyo Night Storm":    "#7AA2F7",
		"Catppuccin Mocha":     "#CBA6F7",
		"Catppuccin Macchiato": "#C6A0F6",
		"Nord":                 "#88C0D0",
		"Gruvbox Dark":         "#FE8019",
		"Gruvbox Light":        "#D65D0E",
		"Kanagawa":             "#7E9CD8",
		"Rose Pine":            "#EBBCBA",
		"Rose Pine Moon":       "#EA9A97",
		"Solarized Dark":       "#268BD2",
		"Solarized Light":      "#268BD2",
		"One Dark":             "#61AFEF",
		"Everforest Dark":      "#A7C080",
		"Monokai":              "#F92672",
	}

	themes := BuiltIns()
	if len(themes) != len(cases) {
		t.Fatalf("brand-accent table covers %d themes but BuiltIns has %d", len(cases), len(themes))
	}

	for _, theme := range themes {
		want, ok := cases[theme.Name]
		if !ok {
			t.Fatalf("no canonical brand accent recorded for theme %q", theme.Name)
		}
		wantColor := lipgloss.Color(want)
		if got := theme.Title.GetForeground(); fmt.Sprint(got) != fmt.Sprint(wantColor) {
			t.Errorf("%s Title foreground: want %v, got %v", theme.Name, wantColor, got)
		}
		if got := theme.HeaderAccent.GetForeground(); fmt.Sprint(got) != fmt.Sprint(wantColor) {
			t.Errorf("%s HeaderAccent foreground: want %v, got %v", theme.Name, wantColor, got)
		}
	}
}

func TestThemeConstructorsDefineCoreStyles(t *testing.T) {
	for _, theme := range BuiltIns() {
		if theme.Text.GetForeground() == nil {
			t.Fatalf("expected %s Text to define a foreground", theme.Name)
		}
		if theme.Accent.GetForeground() == nil {
			t.Fatalf("expected %s Accent to define a foreground", theme.Name)
		}
		if theme.Success.GetForeground() == nil {
			t.Fatalf("expected %s Success to define a foreground", theme.Name)
		}
		if theme.Warn.GetForeground() == nil {
			t.Fatalf("expected %s Warn to define a foreground", theme.Name)
		}
		if theme.Error.GetForeground() == nil {
			t.Fatalf("expected %s Error to define a foreground", theme.Name)
		}
	}
}
