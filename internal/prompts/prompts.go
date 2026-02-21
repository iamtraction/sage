package prompts

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed files/*.md
var fs embed.FS

var templates = template.Must(
	template.New("").
		Option("missingkey=zero").
		ParseFS(fs, "files/*.md"),
)

func Get(prompt string, vars map[string]string) (string, error) {
	tpl := templates.Lookup(prompt + ".md")
	if tpl == nil {
		return "", fmt.Errorf("prompt not found: %s", prompt)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return "", err
	}

	return buf.String(), nil
}
