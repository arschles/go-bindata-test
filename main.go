package main

import (
  "log"
  "net/http"
  "fmt"
)

func main() {
  lorem, err := Asset("static/textfiles/loremipsum.txt")
  if err != nil {
    log.Fatalf(err.Error())
  }

  tmpl, err := GetTemplate("index", "static/templates/index.tmpl")
  if err != nil {
    log.Fatalf(err.Error())
  }

  indexHandler := func(res http.ResponseWriter, req *http.Request) {
    name := req.URL.Query().Get("name")
    if len(name) == 0 {
      name = "go-bindata-test"
    }
    title := req.URL.Query().Get("title")
    if len(title) == 0 {
      title = "Hello World!"
    }
    text := string(lorem)
    tmpl.Execute(res, map[string]string{
      "Title": title,
      "Name": name,
      "Text": text,
    })
  }
  http.HandleFunc("/", indexHandler)

  port := 8080
  log.Printf("listening on %d", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
