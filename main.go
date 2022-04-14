package main

import ("fmt"
        "net/http")

func main() {
  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Home Page")
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Contacts Page")
}
