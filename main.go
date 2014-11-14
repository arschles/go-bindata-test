package main

import (
  "log"
  "html/template"
  "net/http"
  "fmt"
)

func main() {
  tmplBytes, err := Asset("static/templates/index.tmpl")
  if err != nil {
    log.Fatalf(err.Error())
  }
  tmplStr := string(tmplBytes)
  tmpl := template.Must(template.New("index").Parse(tmplStr))

  indexHandler := func(res http.ResponseWriter, req *http.Request) {
    name := req.URL.Query().Get("name")
    if len(name) == 0 {
      name = "go-bindata-test"
    }
    title := req.URL.Query().Get("title")
    if len(title) == 0 {
      title = "Hello World!"
    }
    tmpl.Execute(res, map[string]string{
      "Title": title,
      "Name": name,
    })
  }
  http.HandleFunc("/", indexHandler)

  port := 8080
  log.Printf("listening on %d", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
