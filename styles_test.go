package tuitheme

import "testing"

func TestBuiltInsIncludesPhosphor(t *testing.T) {
	themes := BuiltIns()
	if len(themes) != 1 {
		t.Fatalf("expected one built-in theme, got %d", len(themes))
	}
	if themes[0].Name != "Phosphor" {
		t.Fatalf("expected Phosphor, got %q", themes[0].Name)
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
