package main

import (
  "fmt"
  "net/http"
  "html/template"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
  Id uint16
  Title, Anons, FullText string
}

var posts = []Article{}

func handleFunc() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.HandleFunc("/", index)
  http.HandleFunc("/create", create)
  http.HandleFunc("/save_article", save_article)
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

  readArticles()

  t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

  if (err != nil) {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("title")
  anons := r.FormValue("anons")
  full_text := r.FormValue("full_text")

  if title == "" || anons == "" || full_text == "" {
    fmt.Fprintf(w, "Не все данные заполнены")
  }

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-project")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
  if (err != nil) {
    panic(err)
  }

  defer insert.Close()

  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func readArticles() {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-project")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  res, err := db.Query("SELECT * FROM `articles` ORDER BY id desc")
  if (err != nil) {
    panic(err)
  }

  posts = []Article{}
  for res.Next() {
    var article Article
    err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.FullText)
    if err != nil {
      panic(err)
    }

    posts = append(posts, article)
  }
}
