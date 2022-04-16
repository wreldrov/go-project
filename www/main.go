package main

import (
  "fmt"
  "net/http"
  "html/template"

  "github.com/gorilla/mux"

  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
  Id uint16
  Title, Anons, FullText string
}

var posts = []Article{}
var showPost = Article{}

func handleFunc() {
  r := mux.NewRouter()
  r.HandleFunc("/", index).Methods("GET")
  r.HandleFunc("/create", create).Methods("GET")
  r.HandleFunc("/save_article", save_article).Methods("POST")
  r.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

  http.Handle("/", r)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
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
    var post Article
    err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
    if err != nil {
      panic(err)
    }

    posts = append(posts, post)
  }
}

func readArticle(id string) {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-project")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE id = '%s'", id))
  if (err != nil) {
    panic(err)
  }

  showPost = Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
    if err != nil {
      panic(err)
    }

    showPost = post
  }
}

func show_post(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

  if (err != nil) {
    fmt.Fprintf(w, err.Error())
  }

  vars := mux.Vars(r)
  readArticle(vars["id"])

  t.ExecuteTemplate(w, "show", showPost)
}
