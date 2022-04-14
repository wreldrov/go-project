package main

import ("fmt"; "net/http"; "html/template")

type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grades, Happiness float64
  Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName (newName string) {
  u.Name = newName
}

func main() {
  // var bob User = ....
  // bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 4.2, Happiness: 0.8}

  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
  bob := User{"Bob", 25, 50, 4.2, 0.8, []string{"Football", "Tennis", "Bord"}}
  // fmt.Fprintf(w, `<h1>Home Page</h1>
  //   <b>Main Text</b>`)
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Contacts Page")
}
