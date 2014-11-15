package templatefs

import (
  "html/template"
  "io"
)

type AssetFunc func(string) ([]byte, error)

type Template struct {
  AssetFunc AssetFunc
  tmpl *template.Template
}

func New(name string, fn AssetFunc) *Template {
  return &Template{fn, template.New(name)}
}

func (t *Template) Name() string {
  return t.tmpl.Name()
}

func (t *Template) Funcs(funcMap template.FuncMap) *Template {
  return t.tmpl.Funcs(funcMap)
}

func (t *Template) Parse(filename string) (*Template, error) {
  tmplStr, err := t.file(filename)
  if err != nil {
    return nil, err
  }
  newTmpl, err := t.tmpl.Parse(tmplStr)
  if err != nil {
    return nil, err
  }
  t.tmpl := newTmpl
  return t, nil
}

func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
  fileStrs := []string{}
  for _, filename := range filenames {
    fileStr, err := t.file(filename)
    if err != nil {
      return nil, err
    }
    fileStrs = append(fileStrs, fileStr)
  }
  newTmpl, err := t.tmpl.ParseFiles(fileStrs...)
  if err != nil {
    return nil, err
  }
  t.tmpl = newTmpl
  return t, nil
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
  return t.tmpl.Execute(w, data)
}

func (t *Template) file(name string) ([]byte, error) {
  tmplBytes, err := Asset(fileName)
  if err != nil {
    return nil, err
  }
  return string(tmplBytes), nil
}
