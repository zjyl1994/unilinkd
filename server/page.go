package server

import (
	_ "embed"
	"html/template"
	"io"
	"time"
)

//go:embed tpl.html
var tplData string

//go:embed index.html
var indexData []byte

func writePage(w io.Writer, heading, content string) error {
	tpl, err := template.New("").Parse(tplData)
	if err != nil {
		return err
	}
	return tpl.Execute(w, map[string]interface{}{
		"Title":   heading,
		"Content": content,
		"Time":    time.Now().Format(time.RFC3339),
	})
}
