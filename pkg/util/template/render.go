package template

import (
	"bytes"
	"text/template"
)

// Render reads a template file at the given path and renders it
// with the provided data. It returns the rendered result or an error.
// path: the file path to the template.
// data: the data to populate the template.
func Render(path string, data any) (string, error) {
	tpl, errParse := template.ParseFiles(path)
	if errParse != nil {
		return "", errParse
	}

	var buf bytes.Buffer
	if errExec := tpl.Execute(&buf, data); errExec != nil {
		return "", errExec
	}

	return buf.String(), nil
}
