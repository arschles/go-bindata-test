package templatefs

import (
  "html/template"
)

type AssetFunc func(string) ([]byte, error)

type Template struct {
  AssetFunc AssetFunc
  tmpl *template.Template
}

func New(name string, fn AssetFunc) *Template {
  return &Template{fn, template.New(name)}
}

func (t *Template) Parse(filename string) (*Template, error) {
  tmplBytes, err := Asset(fileName)
  if err != nil {
    return nil, err
  }
  tmplStr := string(tmplBytes)
  newTmpl, err := t.tmpl.Parse(tmplStr)
  if err != nil {
    return nil, err
  }
  t.tmpl := newTmpl
}
