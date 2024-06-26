# {{ .Word }}
## definitions
{{ range .Definitions }}
### {{ .PartOfSpeech }}: {{ .Definition }}

{{ if .Synonyms }}#### Synonyms {{ end }}
{{ range .Synonyms }}
- {{ . }} {{ end }}

{{ if .Antonyms }}#### Antonyms {{ end }}
{{ range .Antonyms }}
- {{ . }} {{ end }}

{{ if .Examples }}#### Examples{{ end }}
{{ range .Examples }}
- {{ . }} {{ end }}

{{ end }}  
## frequency
{{ .Frequency }}
