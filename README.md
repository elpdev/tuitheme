# tuitheme

Small reusable theme primitives for Go terminal user interfaces.

`tuitheme` provides a semantic `Theme` type, a polished default `Phosphor`
theme, and small helpers for selecting themes. Host apps keep ownership of
local themes, command palettes, screen routing, persistence, and branding.

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
