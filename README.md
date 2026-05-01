# tuitheme

Small reusable theme primitives for Go terminal user interfaces.

`tuitheme` provides a semantic `Theme` type, portable built-in themes, and
small helpers for selecting themes. Host apps keep ownership of local themes,
command palettes, screen routing, persistence, and branding.

Built-in themes include Phosphor, Dracula, Tokyo Night, Catppuccin, Nord,
Gruvbox, Kanagawa, Rose Pine, Solarized, One Dark, Everforest, and Monokai
variants.

## Install

```sh
go get github.com/elpdev/tuitheme
```

## Example

```go
func BuiltIns() []tuitheme.Theme {
	return append(tuitheme.BuiltIns(), LocalTheme())
}

func Next(current string) tuitheme.Theme {
	return tuitheme.Next(BuiltIns(), current)
}
```

## Semantic Slots

Themes expose semantic Lip Gloss styles rather than app-specific component
styles. Use the closest semantic slot for each UI state:

- `Text`, `Muted`, `Disabled`
- `Accent`, `Title`, `Selected`, `Focus`
- `Header`, `HeaderAccent`, `Sidebar`, `Main`, `Footer`, `Modal`
- `PaletteAccent`, `Border`
- `Info`, `Success`, `Warn`, `Error`

See `THIRD_PARTY_NOTICES.md` for palette attribution.
