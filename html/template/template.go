package template

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
  return t.replaceTmpl(t.tmpl.Funcs(funcMap))
}

func (t *Template) Parse(filename string) (*Template, error) {
  tmplBytes, err := t.file(filename)
  if err != nil {
    return nil, err
  }
  newTmpl, err := t.tmpl.Parse(string(tmplBytes))
  if err != nil {
    return nil, err
  }
  return t.replaceTmpl(newTmpl), nil
}

func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
  fileBytes := []byte{}
  for _, filename := range filenames {
    tmplBytes, err := t.file(filename)
    if err != nil {
      return nil, err
    }
    fileBytes = append(fileBytes, tmplBytes...)
  }
  newTmpl, err := t.tmpl.Parse(string(fileBytes))
  if err != nil {
    return nil, err
  }
  return t.replaceTmpl(newTmpl), nil
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
  return t.tmpl.Execute(w, data)
}

func (t *Template) replaceTmpl(tmpl *template.Template) *Template {
  t.tmpl = tmpl
  return t
}

func (t *Template) file(fileName string) ([]byte, error) {
  tmplBytes, err := t.AssetFunc(fileName)
  if err != nil {
    return nil, err
  }
  return tmplBytes, nil
}
