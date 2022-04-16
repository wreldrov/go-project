package main

import (
  "fmt"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  fmt.Println("Work with MySQL")

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-project")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  fmt.Println("Connected to MySQL")
}
