package html

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func Render(path string, data any, tpls ...string) (string, error) {
	for i, tpl := range tpls {
		tpls[i] = filepath.Join(path, tpl)
	}
	t, err := template.ParseFiles(tpls...)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
