package main

import (
  "fmt"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id uint16 `json:"id"`
  Name string `json:"name"`
  Age uint16 `json:"age"`
}

func main() {
  fmt.Println("Work with MySQL")

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/go-project")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  // Установка данных
  // insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 35)")
  // if (err != nil) {
  //   panic(err)
  // }
  //
  // defer insert.Close()


  // Выборка данных
  res, err := db.Query("SELECT * FROM `users`")
  if (err != nil) {
    panic(err)
  }

  for res.Next() {
    var user User
    err = res.Scan(&user.Id, &user.Name, &user.Age)
    if err != nil {
      panic(err)
    }

    fmt.Println(fmt.Sprintf("%d. User: %s with age %d", user.Id, user.Name, user.Age))
  }

  fmt.Println("Connected to MySQL")
}
