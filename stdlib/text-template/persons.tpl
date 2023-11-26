# Person Statistics
Count:      {{. | len}}

# Listing
{{/* Minus sign before closing braces means no newlines should be outputed */ -}}
{{range . -}}
---------------------
Name:       {{.Name}}
Age:        {{.Age}}
Gender:     {{.Gender}}
Can Hunt:   {{if .CanHunt}}Yes{{else}}No{{end}}
Skills:     {{range .Skills}}{{.}}, {{else}}Nothing yet!{{end}}
Nick Name:  {{if (eq .Gender "Male")}}Mr {{else}}Miss {{end}}{{.Name}}
=====================
{{end -}}
