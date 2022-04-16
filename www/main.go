package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func handleFunc() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
}

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

  if (err != nil) {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "index", nil)
}
