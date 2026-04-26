# AGENTS.md

Guidance for AI agents and contributors working in this repository.

## Project Intent

`tuitheme` is a reusable theme primitive package for terminal user interfaces.
It provides semantic Lip Gloss style fields, a small default theme set, and tiny
theme selection helpers. Host apps keep ownership of routing, persistence,
settings, command palettes, and app-specific branding.

Keep this package narrowly focused. Do not add app shell behavior or product
features.

## Stack

- Go 1.26+
- Lip Gloss v2 via `charm.land/lipgloss/v2`

## Coding Style

- Prefer small, direct changes.
- Keep exported defaults portable across host apps.
- Default to ASCII in source and docs unless there is a clear reason not to.
- Run `gofmt` on Go files.

## Testing

Before considering work complete, run:

```sh
go test ./...
```
