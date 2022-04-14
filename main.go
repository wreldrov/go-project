package main

import ("fmt"
        "net/http")

func main() {
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil)
}

func home_page(page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(page, "Home Page")
}
