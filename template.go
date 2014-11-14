package main

import (
  "html/template"
)

func GetTemplate(tmplName, fileName string) (*template.Template, error) {
  tmplBytes, err := Asset(fileName)
  if err != nil {
    return nil, err
  }
  return template.New(tmplName).Parse(string(tmplBytes))
}
