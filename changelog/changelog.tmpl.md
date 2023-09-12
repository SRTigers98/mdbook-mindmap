# Release {{ .tag }}
{{ if .pre }}
**This is a pre release! Some functionality might not work as expected!**
{{ end }}
## Features
{{ range .feats }}
- {{ . -}}
{{ end }}

## Fixes
{{ range .fixes }}
- {{ . -}}
{{ end }}
