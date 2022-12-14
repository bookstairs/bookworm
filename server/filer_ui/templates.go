package filer_ui

import (
	_ "embed"
	"html/template"
	"net/url"
	"strings"

	"github.com/dustin/go-humanize"
)

func printpath(parts ...string) string {
	concat := strings.Join(parts, "")
	escaped := url.PathEscape(concat)
	return strings.ReplaceAll(escaped, "%2F", "/")
}

var funcMap = template.FuncMap{
	"humanizeBytes": humanize.Bytes,
	"printpath":     printpath,
}

//go:embed filer.html
var filerHtml string

var StatusTpl = template.Must(template.New("status").Funcs(funcMap).Parse(filerHtml))
