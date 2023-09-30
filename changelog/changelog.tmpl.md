# Release {{ .tag }}
{{ if .pre }}
**This is a pre release! Some functionality might not work as expected!**
{{ end }}
{{- if gt (len .feats) 0}}
## Features
{{ range .feats }}
- {{ . -}}
{{ end }}
{{- end }}
{{- if gt (len .fixes) 0}}
## Fixes
{{ range .fixes }}
- {{ . -}}
{{ end }}
{{- end }}
